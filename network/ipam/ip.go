package ipam

import (
	"encoding/binary"
	"net"
)

var Native binary.ByteOrder = binary.LittleEndian

func ip4Str(arg1 uint32) string {
	ip := make(net.IP, 4)
	Native.PutUint32(ip, arg1)
	return ip.String()
}
