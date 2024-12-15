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

type RemoteEndpointInfo struct {
	SecurityIdentity  uint32 `align:"sec_identity"`
	TunnelEndpoint    uint32 `align:"tunnel_endpoint"`
	TunnelEndpointMac uint64 `align:"tunnel_endpoint_mac"`
	_                 uint16
	_                 uint16
	_                 uint16
	Key               uint8 `align:"key"`
	Flags             uint8 `align:"flag_skip_tunnel"`
}

func EndpointInfoData() {
	g := reflect.TypeOf(RemoteEndpointInfo{})
	bs := binary.Size(RemoteEndpointInfo{})
	rs := int(g.Size())
	_rs := uint32(g.Size())

	fmt.Sprintf("size of EndpointInfo is %d, binary.Size is %d %d", rs, bs, _rs)
}
