package dns

import "net"

func IPv4ToInt(ip net.IP) int64 {
	ipInt := int64(0)
	for _, b := range ip.To4() {
		ipInt = (ipInt << 8) + int64(b)
	}
	return ipInt
}
