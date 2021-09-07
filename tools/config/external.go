package config

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
)

type configChan struct {
	Config *Config
	Error  error
}

type chain struct {
	channel chan *configChan
	client  *clientv3.Client

	key    string
	runEnv runEnv
}

func (c *chain) remote() error {
	conf := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

// local 获取本地的配置文件
func (c *chain) local() (*Config, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(pwd, configFileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	return &config, err
}

// New 要求goroutine启动
func New(key string, env ...runEnv) (*chain, error) {
	rEnv, err := getRunEnv(env)
	if err != nil {
		return nil, err
	}

	return &chain{
		channel: make(chan *configChan, 1),
		key:     key,
		runEnv:  rEnv,
	}, nil
}

// Watch err为初始化或读取或文件反序列化出错，得到channel，便可开始循环读出配置
// 开发模式读取一次后自动结束循环（因为内部close了管道）；生产模式为死循环
func (c *chain) Watch() (<-chan *configChan, error) {
	switch c.runEnv {
	case EnvDev:
		config, err := c.local()
		if err != nil {
			return nil, err
		}

		c.channel <- &configChan{config, nil}
		close(c.channel)

	case EnvProd:
		if err := c.remote(); err != nil {
			return nil, err
		}

		watchChan := c.client.Watch(context.Background(), c.key)
		go func() {
			for res := range watchChan {
				for _, event := range res.Events {
					switch event.Type {
					case mvccpb.PUT:
						var config Config
						err := yaml.Unmarshal(event.Kv.Value, &config)

						c.channel <- &configChan{&config, err}
					}
				}
			}
		}()
	}

	return c.channel, nil
}
