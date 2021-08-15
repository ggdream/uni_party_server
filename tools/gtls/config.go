package gtls

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"strings"
	"time"
)

type childConfig struct {
	Name         string   `yaml:"name"`
	KeyType      KeyType  `yaml:"keyType"`
	Organization string   `yaml:"organization"`
	Duration     string   `yaml:"duration"`
	Hosts        []string `yaml:"hosts"`
}

func (c *childConfig) toOption() (*Options, error) {
	if len(c.Duration) < 2 {
		return nil, errors.New("field duration exist error")
	}
	num, err := strconv.Atoi(c.Duration[:len(c.Duration)-1])
	if err != nil {
		return nil, err
	}

	var duration time.Duration
	if strings.HasSuffix(c.Duration, "y") {
		duration = 365 * 24 * time.Hour * time.Duration(num)
	} else if strings.HasSuffix(c.Duration, "d") {
		duration = 24 * time.Hour * time.Duration(num)
	} else {
		return nil, errors.New("config exist error")
	}

	return &Options{
		Name:         c.Name,
		KeyType:      c.KeyType,
		Organization: c.Organization,
		Duration:     duration,
		Hosts:        c.Hosts,
	}, nil
}

// Config 配置文件结构体
type Config struct {
	Crt      string        `yaml:"crt"`
	Key      string        `yaml:"key"`
	Children []childConfig `yaml:"children"`
}

// CheckFile 检查根证书和根私钥是否存在
func (c *Config) CheckFile() error {
	_, err := os.Stat(c.Crt)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	_, err = os.Stat(c.Key)
	return err
}

// Tran2Options 配置文件字段转 *Options
func (c *Config) Tran2Options() ([]*Options, error) {
	var res []*Options
	for _, child := range c.Children {
		option, err := child.toOption()
		if err != nil {
			return nil, err
		}
		res = append(res, option)
	}

	return res, nil
}

// ParseConfig 解析配置文件
func ParseConfig(name string) (config Config, err error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &config)
	return
}
