package random

import "testing"

func TestNewSMSCode(t *testing.T) {
	for i := 0; i < 6; i++ {
		println(NewSMSCode())
	}
}
