package dataStructure

import "fmt"

func GoArray2C(array []byte, space bool) string {
	ret := ""
	format := ",%#x"
	if space {
		format = ", %#x"
	}

	for i, e := range array {
		if i == 0 {
			ret = ret + fmt.Sprintf("%#x", e)
		} else {
			ret = ret + fmt.Sprintf(format, e)
		}
	}
	return ret
}
