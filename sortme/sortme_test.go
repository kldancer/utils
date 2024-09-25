package sortme

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	a := generate(5)
	fmt.Printf("%v\n", a)
}
