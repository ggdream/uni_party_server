package ggrpc

import "google.golang.org/grpc"

type Client struct {
	conn *grpc.ClientConn
}

// NewClient 实例化一个gRPC客户端
func NewClient() (*Client, error) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

// Conn 获取gRPC的连接句柄
func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

// Close 关闭gRPC连接
func (c *Client) Close() error {
	return c.conn.Close()
}
