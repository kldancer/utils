package network

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/k8snetworkplumbingwg/sriovnet"
	log "github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

var reg1 = regexp.MustCompile(`^([0-9a-fA-F]{4}):([0-9a-fA-F]{4})$`)
var reg2 = regexp.MustCompile(`^([0-9a-fA-F]{4}):([0-9a-fA-F]{2}):([0-9a-fA-F]{2})\.([0-9a-fA-F]{1})$`)

type VfPair struct {
	VfPciAddress string

	// 如果VF不在当前netns中，则为空
	VfIfName string
	// 如果VF Rep不在当前netns中，则为空
	VfRepIfName string
}

// 扫描指定的PF设备，并获取其所有的VF、VF Rep设备信息。
// pf 参数可以传入PF的网卡名称，也可以传入PCI地址，但是必须为完整的地址，如 0000:65:00.0 。
// 返回的 pair 中，vf或者vf rep的名称均有可能为空，为空即该接口不在当前netns中，map的key表示VF的ID。

func ScanVfPairs(pf string) (map[int]*VfPair, error) {
	if strings.Contains(pf, ":") {
		if !reg2.MatchString(pf) {
			return nil, fmt.Errorf("invalid pci address `%s`, must use full address, such as `0000:65:00.0`", pf)
		}

		netdevs, err := sriovnet.GetNetDevicesFromPci(pf)
		if err != nil {
			return nil, fmt.Errorf("can not get net devices from pci: %v", err)
		}

		if len(netdevs) == 0 {
			return nil, fmt.Errorf("no net devices found")
		}

		if len(netdevs) > 1 {
			log.Warnf("more than one net devices found `%v`, use `%s`", netdevs, netdevs[0])
		}

		pf = netdevs[0]
	}

	// 开启SR-IOV，并设置最大的VF
	err := sriovnet.EnableSriov(pf)
	if err != nil {
		return nil, fmt.Errorf("can not enable sriov: %v", err)
	}

	// 获取PF的handle，其实也就是一个列表，仅能在当前流程中使用
	handle, err := sriovnet.GetPfNetdevHandle(pf)
	if err != nil {
		return nil, fmt.Errorf("can not get pf netdev handle: %v", err)
	}

	// 从PF设备获取PCI地址
	pfPci, err := sriovnet.GetPciFromNetDevice(pf)
	if err != nil {
		return nil, fmt.Errorf("can not get pci from net device: %v", err)
	}

	// 通过PF设备的PCI地址获取它的面板口的Rep口
	uplink, err := sriovnet.GetUplinkRepresentor(pfPci)
	if err != nil {
		return nil, fmt.Errorf("can not get uplink representor: %v", err)
	}

	vfPairs := make(map[int]*VfPair, len(handle.List))
	for _, vfObj := range handle.List {
		vf := sriovnet.GetVfNetdevName(handle, vfObj)

		// 用面板口Rep找VF的Rep口
		vfRep, err := sriovnet.GetVfRepresentor(uplink, vfObj.Index)
		if err != nil {
			log.Debugf("can not get vf representor with pf `%s` and vf pci `%s`: %v, skipping...",
				pf, vfObj.PciAddress, err)
			continue
		}

		vfPairs[vfObj.Index] = &VfPair{
			VfPciAddress: vfObj.PciAddress,
			VfIfName:     vf,
			VfRepIfName:  vfRep,
		}
	}

	return vfPairs, nil
}

// 预先修复已有的VF Pairs。
// 主要是针对已经存在的接口，检查是否有一端在用，如果都没在用，就检查名称是否为fab开头，如果是就改掉，避免和新建的pod冲突。
func prefixVfPairs(vfPairs map[int]*VfPair) error {
	for id, pair := range vfPairs {
		// 当VF在容器里（即在使用中）时，或者VF Rep出现问题返回空值时，都跳过
		if pair.VfIfName == "" || pair.VfRepIfName == "" {
			continue
		}

		// 如果有VF口没有被attach成功，但是仍然有fab开头名称的，那么就把它们改名
		if strings.HasPrefix(pair.VfRepIfName, "fab") {
			link, err := netlink.LinkByName(pair.VfRepIfName)
			if err != nil {
				return fmt.Errorf("can not get link by name `%s`: %v", pair.VfRepIfName, err)
			}

			err = netlink.LinkSetDown(link)
			if err != nil {
				return fmt.Errorf("can not set link down: %v", err)
			}

			err = netlink.LinkSetName(link, fmt.Sprintf("fab-vf_rep-%d", id))
			if err != nil {
				return fmt.Errorf("can not set link name: %v", err)
			}
		}
	}

	return nil
}

// 使用vendorId、deviceId扫描PCI设备，返回所有匹配的PCI地址。
// 参数格式为 `vendorId:deviceId`，如 `8086:10ed`。

func ScanPciDevice(vendorDeviceId string) ([]string, error) {
	if !reg1.MatchString(vendorDeviceId) {
		return nil, fmt.Errorf("invalid vendor/device id `%s`, must use full address, such as `8086:10ed`", vendorDeviceId)
	}

	args := strings.Split(vendorDeviceId, ":")

	pci, err := ghw.PCI()
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0)
	for _, dev := range pci.Devices {
		if dev.Vendor.ID == args[0] && dev.Product.ID == args[1] {
			ret = append(ret, dev.Address)
		}
	}

	return ret, nil
}
