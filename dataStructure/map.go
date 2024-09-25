package dataStructure

import "fmt"

func MapFor() {

	aMap := make(map[string]string)
	aMap["key1"] = "value1"
	aMap["key2"] = "value2"

	for k, v := range aMap {
		v += "123"
		aMap[k] = v
	}

	fmt.Printf("map: %v\n", aMap)

}

func Math() {
	a := 13 * 5
	b := 12
	c := a / b
	d := 13 * 5 / 12

	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
}
