package dataStructure

import (
	"fmt"
	"testing"
)

func TestGoArray2C(t *testing.T) {
	array := []byte{0x00, 0x0c, 0x29, 0x1e, 0x7f, 0x66}
	ret := GoArray2C(array, true)

	fmt.Print(ret)
}
