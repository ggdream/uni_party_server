package ggrpc

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_Conn(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	newClient := NewGreetClient(client.Conn())
	res, err := newClient.Say(context.TODO(), &Request{Name: "wang"})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Res)
}
