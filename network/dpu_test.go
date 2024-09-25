package network

import (
	"fmt"
	"log"
	"testing"
)

func TestScanPciDevice(t *testing.T) {
	devs, err := ScanPciDevice("1da8:4000")
	if err != nil {
		panic(err)
	}

	fmt.Printf("devs: %v\n", devs)
}

func TestScanVfPairs(t *testing.T) {
	pairs, err := ScanVfPairs("0000:05:00.0")
	if err != nil {
		log.Fatal(err)
	}
	pairs2, err := ScanVfPairs("0000:0b:00.0")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(pairs); i++ {
		fmt.Printf("  - VfId \t`%d`: \t\tVf \t`%s`, \t\tVf Rep \t`%s`\n",
			i, pairs[i].VfIfName, pairs[i].VfRepIfName)
	}
	fmt.Println("-------------------")
	for i := 0; i < len(pairs2); i++ {
		fmt.Printf("  - VfId \t`%d`: \t\tVf \t`%s`, \t\tVf Rep \t`%s`\n",
			i, pairs2[i].VfIfName, pairs2[i].VfRepIfName)
	}
}
