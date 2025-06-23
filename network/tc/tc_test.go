package tc

import (
	"testing"
)

func TestAddIfTcFilter(t *testing.T) {
	if err := AddIfTcFilter("enp3s0", "10.0.112.0/24", []int{22, 23, 24}); err != nil {
		t.Errorf("AddIfTcFilter failed: %v", err)
	}
}

func TestDeleteIfTcFilter(t *testing.T) {
	if err := deleteIfTcFilter("enp3s0"); err != nil {
		t.Errorf("failed to delete tc filter: %v", err)
	}
}

func TestAdd2(t *testing.T) {
	add2()
}
