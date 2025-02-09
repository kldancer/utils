package ipam

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	ips := ip4Str(234946571)
	fmt.Sprintf("%s", ips)
}
