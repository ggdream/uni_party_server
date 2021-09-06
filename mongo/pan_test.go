package mongo

import (
	"fmt"
	"testing"
	"time"
)

func TestN(t *testing.T) {
	value := []int{1, 2, 3}
	for _, v := range value {
		fmt.Println(value)
		value = append(value[1:], v*2)
		time.Sleep(1*time.Second)
	}
}
