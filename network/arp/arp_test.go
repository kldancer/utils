package arp

import (
	"fmt"
	"testing"
)

func TestArp(t *testing.T) {
	mac, err := GetMACAddressByIP("192.168.31.130")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("MAC Address for IP %v is %v\n", "ip", mac)
}
