package system

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/google/martian/log"
	"golang.org/x/crypto/ssh"
)

func GetCpuUsedInHost(hostIp, user, paasword, port string, testTime int) (float64, error) {

	// 通过 SSH 连接
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(paasword),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", hostIp+":"+port, sshConfig)
	if err != nil {
		return 0, fmt.Errorf("cannot connect to host %s, err:%s", hostIp, err)
	}
	defer sshClient.Close()

	// 在 server2 上执行命令获取 CPU 利用率
	session, err := sshClient.NewSession()
	if err != nil {
		return 0, fmt.Errorf("cannot create session for host %s, err:%s", hostIp, err)
	}
	defer session.Close()

	cmd := fmt.Sprintf("sar -u 1 %d|tail -n 1|awk '{print $3,$5}'", testTime)
	log.Infof("cmd: %s", cmd)
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return 0, fmt.Errorf("cannot get cpu usage for host %s, err:%s", hostIp, err)
	}
	var cpuUsage float64
	CpuUseList := strings.Split(string(output), " ")
	for _, s := range CpuUseList {
		usage, err := strconv.ParseFloat(strings.TrimRight(s, "\n"), 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse cpu usage for host %s, err:%s", hostIp, err)
		}
		cpuUsage += usage
	}
	return cpuUsage, nil
}

func GetInterfaceMac(name string) (string, error) {
	// 根据接口名称查找接口
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}

	// 获取接口的 MAC 地址
	macAddr := iface.HardwareAddr

	return macAddr.String(), err
}
