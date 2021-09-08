package etcd

import (
	"context"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/client/v3"
	"time"
)

var client *EtcD

func Init() (err error) {
	client, err = New()
	return
}

type EtcD struct {
	client  *clientv3.Client
	context context.Context
}

func New() (*EtcD, error) {
	conf := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	_, err = client.UserGet(ctx, "test")
	if err != nil {
		return nil, errors.New("cannot connect the etcd server")
	}

	return &EtcD{
		client:  client,
		context: context.Background(),
	}, nil
}
