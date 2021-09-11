package ggrpc

import (
	"google.golang.org/grpc"
	"net"
)

// Server server端封装体
type Server struct {
	server *grpc.Server
}

// NewServer 实例化一个gRPC服务
func NewServer(register func(grpc.ServiceRegistrar)) *Server {
	server := grpc.NewServer()
	register(server)

	return &Server{server: server}
}

// Serve 运行服务
func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

// Stop 停止服务
func (s *Server) Stop() {
	s.server.Stop()
}
