package network

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/vishvananda/netlink"
)

func BondAdd(bondName string) {
	// 创建一个bond接口
	bond := &netlink.Bond{
		LinkAttrs: netlink.LinkAttrs{
			Name: bondName,
		},
		Mode:           netlink.BOND_MODE_BALANCE_XOR,
		XmitHashPolicy: netlink.BOND_XMIT_HASH_POLICY_LAYER3_4,
	}

	// 添加bond接口
	if err := netlink.LinkAdd(bond); err != nil {
		fmt.Printf("无法添加bond接口: %v\n", err)
		return
	}

	fmt.Println("成功创建并配置了bond接口:", bondName)
}

func GetBondSlaves(bondName string) ([]string, error) {
	bond, err := netlink.LinkByName(bondName)
	if err != nil {
		return nil, err
	}

	slaves, err := netlink.LinkList()
	if err != nil {
		return nil, err
	}

	var slaveNames []string
	for _, slave := range slaves {
		if slave.Attrs().MasterIndex == bond.Attrs().Index {
			slaveNames = append(slaveNames, slave.Attrs().Name)
		}
	}

	return slaveNames, nil
}

func GetEthtoolLine(interfaceName string) {

	cmd := exec.Command("ethtool", "-i", interfaceName)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	driverLine := ""
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "driver:") {
			driverLine = line
			break
		}
	}

	if driverLine == "" {
		fmt.Println("Driver information not found")
		return
	}

	driverInfo := strings.TrimSpace(strings.TrimPrefix(driverLine, "driver:"))
	fmt.Println("Driver:", driverInfo)
}
