package uuid

import "testing"

func TestNewV1(t *testing.T) {
	println(NewV1())
}

func TestNewV4(t *testing.T) {
	println(NewV4())
}

func TestNew(t *testing.T) {
	println(New())
}
