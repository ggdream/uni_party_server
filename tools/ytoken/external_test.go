package ytoken

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	handler := Handler{key: []byte("0123456789")}
	token, err := handler.Sign(&YToken{Protocol: DefaultProtocol, Constraint: &Constraint{
		Signer: "mocaraka",
		Expiry: time.Now().Add(time.Hour).Unix(),
		Serial: "102",
	}})
	if err != nil {
		panic(err)
	}
	println(token)

	yToken, isEqual, err := handler.Verify(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(isEqual)
	fmt.Printf("%v\n", yToken)
	fmt.Println(yToken.Timeout())
}
