package ggrpc

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

type value struct {
	UnimplementedGreetServer
}

func (v *value) Say(ctx context.Context, req *Request) (*Response, error) {
	return &Response{
		Res: req.Name + ", hello",
	}, nil
}

func Fu(r grpc.ServiceRegistrar) {
	RegisterGreetServer(r, new(value))
}

func TestServer_Serve(t *testing.T) {
	serve := NewServer(Fu)
	if err := serve.Serve(); err != nil {
		panic(err)
	}
}
