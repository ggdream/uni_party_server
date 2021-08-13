package random

import "testing"

func TestNewDeviceID(t *testing.T) {
	for i := 0; i < 6; i++ {
		println(NewDeviceID())
	}
}
