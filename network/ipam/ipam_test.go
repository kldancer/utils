package ipam

import (
	"math/big"
	"net"
	"testing"
)

func TestParseAvailableNetwork(t *testing.T) {
	network := &Network{
		CIDRBlock: "10.245.0.0/16",
		Gateway:   "",
		Mask:      0,
	}

	blockCIDR, networkIP, broadcastIP := ParseAvailableNetwork(network)
	t.Log(blockCIDR, networkIP, broadcastIP)
}

func TestRangeSize(t *testing.T) {
	_, cidr, err := net.ParseCIDR("192.168.1.0/24")
	size := RangeSize(cidr)
	t.Logf("size: %d, err: %v", size, err)
	base := bigForIP(cidr.IP)
	t.Logf("base: %d", base)
}

func TestAllocateBit(t *testing.T) {
	allo := big.NewInt(0)
	at, ok := AllocateBit(allo, 254, 0)
	t.Logf("at: %d, ok: %v", at, ok)
}

func TestAddIPOffset(t *testing.T) {
	_, cidr, _ := net.ParseCIDR("192.168.1.0/24")
	base := bigForIP(cidr.IP)
	ip := AddIPOffset(base, 2)
	t.Logf("ip: %s", ip)
}

func TestCalculateIPOffset(t *testing.T) {
	_, cidr, _ := net.ParseCIDR("192.168.1.0/24")
	base := bigForIP(cidr.IP)
	ip := net.IP{192, 168, 1, 3}
	t.Logf("offset: %d", CalculateIPOffset(base, ip))
}
