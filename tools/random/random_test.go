package random

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	for i := 0; i < 6; i++ {
		fmt.Println(New(32))
	}
}

func TestDefault(t *testing.T) {
	for i := 0; i < 6; i++ {
		fmt.Println(Default())
	}
}
