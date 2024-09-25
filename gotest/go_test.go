package gotest

import "testing"

func TestGo(t *testing.T) {
	Go1()
}

func TestStringsSplit(t *testing.T) {
	StringsSplit()
}

func TestIfOr(t *testing.T) {
	IfOr()
}

func TestIntRand(t *testing.T) {
	intRand(12312321.5123341223123)
}

func TestTerror(t *testing.T) {
	terror()
}

func TestRandomMac(t *testing.T) {
	str := randomMac()
	t.Log(str)
}
