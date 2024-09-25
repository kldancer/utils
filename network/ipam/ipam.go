package ipam

import (
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"sync"
)

type Network struct {
	CIDRBlock string
	Gateway   string
	Mask      int
}

func ParseAvailableNetwork(network *Network) (blockCIDR *net.IPNet, networkIP, broadcastIP string) {
	var networkCIDR *net.IPNet
	_, blockCIDR, _ = net.ParseCIDR(network.CIDRBlock)

	if network.Mask > 0 && network.Gateway != "" {
		_, networkCIDR, _ = net.ParseCIDR(fmt.Sprintf("%s/%d", network.Gateway, network.Mask))
	} else {
		_, networkCIDR, _ = net.ParseCIDR(network.CIDRBlock)
	}

	networkIP = networkCIDR.IP.String()
	for i := 0; i < len(networkCIDR.IP); i++ {
		networkCIDR.IP[i] = networkCIDR.IP[i] | (^networkCIDR.Mask[i])
	}
	broadcastIP = networkCIDR.IP.String()

	return blockCIDR, networkIP, broadcastIP
}

// AllocationBitmap is a contiguous block of resources that can be allocated atomically.
//
// Each resource has an offset.  The internal structure is a bitmap, with a bit for each offset.
//
// If a resource is taken, the bit at that offset is set to one.
// r.count is always equal to the number of set bits and can be recalculated at any time
// by counting the set bits in r.allocated.
//
// TODO: use RLE and compact the allocator to minimize space.
type AllocationBitmap struct {
	// strategy carries the details of how to choose the next available item out of the range
	strategy bitAllocator
	// max is the maximum size of the usable items in the range
	max int
	// rangeSpec is the range specifier, matching RangeAllocation.Range
	rangeSpec string

	// lock guards the following members
	lock sync.Mutex
	// count is the number of currently allocated elements in the range
	count int
	// allocated is a bit array of the allocated items in the range
	allocated *big.Int
}

// bitAllocator represents a search strategy in the allocation map for a valid item.
type bitAllocator interface {
	AllocateBit(allocated *big.Int, max, count int) (int, bool)
}

// Range is a contiguous block of IPs that can be allocated atomically.
//
// The internal structure of the range is:
//
//	For CIDR 10.0.0.0/24
//	254 addresses usable out of 256 total (minus base and broadcast IPs)
//	  The number of usable addresses is r.max
//
//	CIDR base IP          CIDR broadcast IP
//	10.0.0.0                     10.0.0.255
//	|                                     |
//	0 1 2 3 4 5 ...         ... 253 254 255
//	  |                              |
//	r.base                     r.base + r.max
//	  |                              |
//	offset #0 of r.allocated   last offset of r.allocated
type Range struct {
	net *net.IPNet
	// base is a cached version of the start IP in the CIDR range as a *big.Int
	base *big.Int
	// max is the maximum size of the usable addresses in the range
	max int

	alloc Interface
}

// RangeSize returns the size of a range in valid addresses.
func RangeSize(subnet *net.IPNet) int64 {
	ones, bits := subnet.Mask.Size()
	if bits == 32 && (bits-ones) >= 31 || bits == 128 && (bits-ones) >= 127 {
		return 0
	}
	// For IPv6, the max size will be limited to 65536
	// This is due to the allocator keeping track of all the
	// allocated IP's in a bitmap. This will keep the size of
	// the bitmap to 64k.
	if bits == 128 && (bits-ones) >= 16 {
		return int64(1) << uint(16)
	} else {
		return int64(1) << uint(bits-ones)
	}
}

// bigForIP creates a big.Int based on the provided net.IP
func bigForIP(ip net.IP) *big.Int {
	b := ip.To4()
	if b == nil {
		b = ip.To16()
	}
	return big.NewInt(0).SetBytes(b)
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AllocateBit(allocated *big.Int, max, count int) (int, bool) {
	if count >= max {
		return 0, false
	}
	offset := rand.Intn(max)
	for i := 0; i < max; i++ {
		at := (offset + i) % max
		if allocated.Bit(at) == 0 {
			return at, true
		}
	}
	return 0, false
}

// addIPOffset adds the provided integer offset to a base big.Int representing a
// net.IP
func AddIPOffset(base *big.Int, offset int) net.IP {
	return net.IP(big.NewInt(0).Add(base, big.NewInt(int64(offset))).Bytes())
}

// calculateIPOffset calculates the integer offset of ip from base such that
// base + offset = ip. It requires ip >= base.
func CalculateIPOffset(base *big.Int, ip net.IP) int {
	ipInt := bigForIP(ip)
	a := big.NewInt(0).Sub(ipInt, base)

	return int(a.Int64())
}
