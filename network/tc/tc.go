package tc

import (
	"fmt"
	"os/exec"
	"strconv"
)

func AddIfTcFilter(ifName string, cidr string, ports []int) error {
	for _, p := range ports {
		if err := addTcFilter(ifName, cidr, strconv.Itoa(p), true); err != nil {
			return fmt.Errorf("faile to add tc filter ifName %s src port %d : %v", ifName, p, err)
		}
		if err := addTcFilter(ifName, cidr, strconv.Itoa(p), false); err != nil {
			return fmt.Errorf("faile to add tc filter ifName %s dst port %d : %v", ifName, p, err)
		}
	}
	return nil
}

func addTcFilter(ifName string, cidr string, port string, isSrc bool) error {
	args := []string{
		"filter", "add", "dev", ifName, "parent", "1:", "protocol", "ip", "prio", "1", "flower",
		"ip_proto", "tcp",
		"src_ip", cidr,
		"src_port", port,
		"action", "skbedit", "priority", "1", "flowid", "1:2",
	}

	if !isSrc {
		args[13] = "dst_ip"
		args[15] = "dst_port"
	}

	_, err := exec.Command("tc", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("faile to add tc filter: %v", err)
	}
	return nil
}

func deleteIfTcFilter(ifName string) error {
	_, err := exec.Command("tc", "filter", "del", "dev", ifName).CombinedOutput()
	if err != nil {
		return fmt.Errorf("faile to delete tc filter: %v", err)
	}
	return nil
}

func htons(val uint16) uint16 {
	return (val<<8)&0xff00 | val>>8
}

func uint16Ptr(val uint16) *uint16 {
	return &val
}

func add2() {
	cmd := exec.Command("tc", "filter", "add", "dev", "enp3s0", "parent", "1:", "protocol", "ip", "prio", "1", "flower",
		"ip_proto", "tcp",
		"src_ip", "10.0.112.0/24",
		"src_port", "22",
		"action", "skbedit", "priority", "1", "flowid", "1:2")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\nOutput: %s\n", err, output)
		return
	}

	fmt.Println("Command executed successfully.")
}
