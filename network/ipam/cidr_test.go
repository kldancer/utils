package ipam

import (
	"fmt"
	"testing"
)

func TestIsIPInSubnet(t *testing.T) {
	ip := "192.168.1.10"
	subnet := "192.168.1.0"
	maskBits := 24

	if isIPInSubnet(ip, subnet, maskBits) {
		fmt.Println("IP is in the subnet")
	} else {
		fmt.Println("IP is not in the subnet")
	}
}

func TestCArrayStringForIPv4(t *testing.T) {
	ip := "10.210.20.152"
	a := CArrayStringForIPv4(ip)
	fmt.Println(a)

	_a, _ := ipv4ToCFormat(ip)
	fmt.Println(_a)

}
