package gotest

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

type ClusterServiceName struct {
	Id string `json:"id"` // cluster id
	Ns string `json:"ns"` // namespace
	Nm string `json:"nm"` // name
}

func ServiceNameFromLeader(name string) ClusterServiceName {
	var serviceName ClusterServiceName
	json.Unmarshal([]byte(name), &serviceName)
	return serviceName
}

func ServiceNameToLeader(serviceName ClusterServiceName) string {
	name, _ := json.Marshal(serviceName)
	return string(name)
}

func Go() {

	c := ClusterServiceName{
		Id: "cluster-1",
		Ns: "ns-test",
		Nm: "svc-1",
	}

	name := ServiceNameToLeader(c)
	fmt.Println("ServiceNameToLeader = ", name)
	s := ServiceNameFromLeader(name)
	fmt.Printf("ServiceNameFromLeader = %+v\n", s)
	//fmt.Println("ServiceNameToLeader1(c) = ", ServiceNameToLeader1(c))
	return
}

func ServiceNameFromLeader1(name string) ClusterServiceName {
	nameStr := strings.Split(name, ".")
	return ClusterServiceName{
		Id: nameStr[0],
		Ns: nameStr[1],
		Nm: nameStr[2],
	}
}

func ServiceNameToLeader1(serviceName ClusterServiceName) string {
	name := serviceName.Id + "." + serviceName.Ns + "." + serviceName.Nm
	return name
}

func Go1() {

	c := ClusterServiceName{
		Id: "cluster-1",
		Ns: "ns-test",
		Nm: "svc-1",
	}

	name := ServiceNameToLeader1(c)
	fmt.Println("ServiceNameToLeader = ", name)
	s := ServiceNameFromLeader1(name)
	fmt.Printf("ServiceNameFromLeader = %+v\n", s)
	//fmt.Println("ServiceNameToLeader1(c) = ", ServiceNameToLeader1(c))
	return
}

func StringsSplit() {
	str1 := "123asBdf"
	str2 := "asdfG"

	result1 := strings.Split(str1, "G")
	fmt.Printf("result1 = %+v\n", result1)
	fmt.Printf("result1.length = %+v\n", len(result1))

	result2 := strings.Split(str2, "G")
	fmt.Printf("result2 = %+v\n", result2)
	fmt.Printf("result2.length = %+v\n", len(result2))
}

func IfOr() {
	a := "None"
	if len(a) == 0 || a == "None" {
		fmt.Println("a")
	}

	b := ""
	if len(b) == 0 || b == "None" {
		fmt.Println("b")
	}

	c := "123"
	if len(c) == 0 || c == "None" {
		fmt.Println("c")
	}

	var d string
	if len(d) == 0 || d == "None" {
		fmt.Println("d")
	}
}

func intRand(cpuUsage float64) {
	fmt.Printf("result = %.5f", cpuUsage)
}

func b() error {
	return fmt.Errorf("b error")
}
func a() error {
	if err := b(); err != nil {
		return fmt.Errorf("a call b error: %v", err)
	}

	return fmt.Errorf("a error")
}
func terror() {
	if err := a(); err != nil {
		log.Errorf("terror: %#-v", err)
	}
}

func randomMac() string {
	r := rand.Uint64()
	defer runtime.KeepAlive(r)

	b := *(*[8]byte)(unsafe.Pointer(&r))

	b[2] = 0xFA
	b[3] = 0xB0
	b[4] = 0x00

	_, _ = fmt.Sscanf("FA:B0:00", "%x:%x:%x", &b[2], &b[3], &b[4])

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", b[2], b[3], b[4], b[5], b[6], b[7])
}
