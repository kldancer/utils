package arp

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"log"
	"net"
	"net/netip"
	"time"

	"github.com/mdlayher/arp"
)

func GetMac() {
	// 替换为你的网络接口名称（例如 "eth0"）
	iface, err := net.InterfaceByName("en0")
	if err != nil {
		log.Fatalf("Failed to get interface: %v", err)
	}

	// 替换为目标 IP 地址
	addr, err := netip.ParseAddr("192.168.31.130")
	if err != nil {
		log.Fatalf("Invalid IP address: %v", err)
	}

	// 打开 ARP 客户端
	c, err := arp.Dial(iface)
	if err != nil {
		log.Fatalf("Failed to dial ARP: %v", err)
	}
	defer c.Close()

	// 设置超时
	c.SetDeadline(time.Now().Add(3 * time.Second))

	// 发送请求并获取 MAC 地址
	mac, err := c.Resolve(addr)
	if err != nil {
		log.Fatalf("Failed to resolve MAC address: %v", err)
	}

	fmt.Printf("MAC address of %s: %s\n", addr, mac)
}

func GetMACAddressByIP(ipStr string) (net.HardwareAddr, error) {
	ip := net.ParseIP(ipStr)
	addrs, err := netlink.NeighList(0, 0)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if addr.IP.Equal(ip) {
			return addr.HardwareAddr, nil
		}
	}
	return nil, fmt.Errorf("no ARP entry found for IP %v", ip)
}
