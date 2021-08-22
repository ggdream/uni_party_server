package group

import (
	"fmt"
	"testing"
)

type test struct {
	Data	*[]string
}

func TestName(t *testing.T) {
	data := make([]string, 0)
	te := test{&data}
	for i := 0; i < 100000; i++ {
		data = append(data, string(rune(i)))
	}
	fmt.Println(te.Data)
}
