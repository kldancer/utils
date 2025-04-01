package ipam

import (
	"fmt"
	"net"
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

func CArrayStringForIPv4(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil || ip.To4() == nil {
		return "0"
	}

	ipv4 := ip.To4()
	ipInt := uint32(ipv4[0])<<24 | uint32(ipv4[1])<<16 | uint32(ipv4[2])<<8 | uint32(ipv4[3])
	return fmt.Sprintf("%d", ipInt)
}

func ipv4ToCFormat(ip string) (string, error) {
	// 解析IPv4地址
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", fmt.Errorf("invalid IP address")
	}

	// 如果是IPv6地址，取其IPv4部分
	if parsedIP.To4() == nil {
		return "", fmt.Errorf("not an IPv4 address")
	}

	// 转换为32位无符号整数
	ipv4 := parsedIP.To4()
	ipUint32 := uint32(ipv4[0])<<24 | uint32(ipv4[1])<<16 | uint32(ipv4[2])<<8 | uint32(ipv4[3])

	// 生成宏定义形式
	return fmt.Sprintf("（0x%X, 0x%X, 0x%X, 0x%X)",
		(ipUint32>>24)&0xFF,
		(ipUint32>>16)&0xFF,
		(ipUint32>>8)&0xFF,
		ipUint32&0xFF), nil
}
func uint32ToIpv4(ipUint32 uint32) string {
	// 提取每个字节并格式化为字符串
	return fmt.Sprintf("%d.%d.%d.%d",
		(ipUint32>>24)&0xFF, // 提取最高字节
		(ipUint32>>16)&0xFF, // 提取次高字节
		(ipUint32>>8)&0xFF,  // 提取次低字节
		ipUint32&0xFF)       // 提取最低字节
}
