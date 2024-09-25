package dataStructure

import (
	"fmt"
	"math/big"
)

func BigIntTest() {
	// 创建一个大整数对象
	num := new(big.Int)
	num.SetString("1234567890123456789012345678901234567890", 10)

	// 获取第 5 位的值
	bitValue := num.Bit(4)

	fmt.Printf("Bit 4 value: %d\n", bitValue)
}
