package ipam

import (
	"strconv"
	"strings"
)

func ipToUint32(ip string) uint32 {
	parts := strings.Split(ip, ".")
	var ipNum uint32
	for i := 0; i < 4; i++ {
		// 通过将 ipNum 的二进制形式左移8个比特位操作，我们为IP地址的下一个字节腾出了位置，然后将当前字节的值加到 ipNum 中。
		part, _ := strconv.Atoi(parts[i])
		ipNum = ipNum<<8 + uint32(part)
	}
	return ipNum
}

func isIPInSubnet(ip, subnet string, maskBits int) bool {
	ipNum := ipToUint32(ip)
	subnetNum := ipToUint32(subnet)
	// ^uint32(0) 生成一个所有位都为 1 的32位无符号整数
	// << (32 - maskBits) 将该整数左移 32 - maskBits 位，相当于生成了一个子网掩码
	mask := ^uint32(0) << (32 - maskBits)

	// 计算IP地址在子网掩码下的网络部分、 计算子网地址在子网掩码下的网络部分。如果相同就是属于该子网
	return (ipNum & mask) == (subnetNum & mask)
}
