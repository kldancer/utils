package dns

import (
	"fmt"
	"net"
	"testing"
)

func TestIPv4ToInt(t *testing.T) {
	a := IPv4ToInt(net.ParseIP("172.28.100.36"))
	fmt.Printf("%d", a)
}
