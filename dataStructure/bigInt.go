package dataStructure

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
)

func BigIntTest() {
	// 创建一个大整数对象
	num := new(big.Int)
	num.SetString("1234567890123456789012345678901234567890", 10)

	// 获取第 5 位的值
	bitValue := num.Bit(4)

	fmt.Printf("Bit 4 value: %d\n", bitValue)
}

type EndpointInfo struct {
	IfIndex  uint32 `align:"ifindex"`
	Unused   uint16 `align:"unused"`
	LxcID    uint16 `align:"lxc_id"`
	Flags    uint32 `align:"flags"`
	VpcFlags uint16 `align:"vpc_flags"`
	// go alignment
	//_       uint64
	MAC     uint64    `align:"mac"`
	NodeMAC uint64    `align:"node_mac"`
	SecID   uint32    `align:"sec_id"`
	Pad     [3]uint32 `align:"pad"`
}

// 结构体整体大小必须是其最大字段大小（8 字节）的倍数。

type RemoteEndpointInfo struct {
	SecurityIdentity  uint32 `align:"sec_identity"`
	TunnelEndpoint    uint32 `align:"tunnel_endpoint"`
	TunnelEndpointMac uint64 `align:"tunnel_endpoint_mac"`
	_                 uint32
	_                 uint16
	//_                 uint16
	Key   uint8 `align:"key"`
	Flags uint8 `align:"flag_skip_tunnel"`
}

type EndpointKey struct {
	// represents both IPv6 and IPv4 (in the lowest four bytes)
	IP        [16]byte `align:"$union0"`
	Family    uint8    `align:"family"`
	Key       uint8    `align:"key"`
	ClusterID uint16   `align:"cluster_id"`
}

func EndpointInfoData() {
	g := reflect.TypeOf(RemoteEndpointInfo{})
	bs := binary.Size(RemoteEndpointInfo{})
	rs := int(g.Size())
	_rs := uint32(g.Size())

	fmt.Sprintf("size of EndpointInfo is %d, binary.Size is %d %d", rs, bs, _rs)
}

type AffixKey struct {
	Family uint8    `align:"family"`
	IP     [16]byte `align:"$union0"`
}

func AffixKeyData() {
	g := reflect.TypeOf(AffixKey{})
	bs := binary.Size(AffixKey{})
	rs := int(g.Size())
	_rs := uint32(g.Size())

	fmt.Sprintf("size of AffixKey is %d, binary.Size is %d %d", rs, bs, _rs)
}
