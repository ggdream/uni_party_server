package mongo

import (
	"testing"
)

func TestInit(t *testing.T) {
	if err := Init(); err != nil {
		panic(err)
	}
}
