package network

import (
	"fmt"
	"testing"
)

func TestBondAdd(t *testing.T) {
	BondAdd("bond0")
}

func TestGetBondSlaves(t *testing.T) {
	slaves, err := GetBondSlaves("bond-test")
	if err != nil {
		fmt.Printf("Error getting bond slaves: %v\n", err)
		return
	}

	fmt.Printf("Slaves of bond0: %v\n", slaves)
}

func TestGetEthtoolLine(t *testing.T) {
	GetEthtoolLine("eth0")
}
