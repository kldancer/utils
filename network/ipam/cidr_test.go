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
