package system

import (
	"fmt"
	"testing"
)

func TestGetCpuUsedInHost(t *testing.T) {

	cpuUse, err := GetCpuUsedInHost("118.31.62.83", "root", "Ykl424713", "22", 2)
	if err != nil {
		fmt.Errorf("err:%s", err)
	}
	fmt.Println("cpuUse:", cpuUse)
}
