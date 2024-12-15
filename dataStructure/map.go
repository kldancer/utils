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

func mapWork() {
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 1
	a["c"] = 0

	b := mapWorker(a)
	fmt.Printf("b: %v\n", b)
}

func mapWorker(m map[string]int) map[string]int {
	m["b"] += 1
	m["c"] += 1
	return m
}
