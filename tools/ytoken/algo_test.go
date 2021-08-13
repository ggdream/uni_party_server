package ytoken

import (
	"testing"
)

func TestAlgoType_Func(t *testing.T) {
	algo := HS256.Func([]byte("test"))

	message := "我喜欢你"
	res := algo.Sign([]byte(message))
	println(algo.Verify([]byte(message), res))
}
