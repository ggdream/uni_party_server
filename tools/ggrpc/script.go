package ggrpc

//go:generate protoc --go_out=. --go-grpc_out=. test.proto

//go:generate protoc --gogo_out=. --go-grpc_out=. test.proto

//go:generate protoc --gofast_out=. --go-grpc_out=. test.proto
