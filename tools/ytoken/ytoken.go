package ytoken

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
	"time"
	"unsafe"
)

var (
	DefaultProtocol = &Protocol{
		Type: "tk",
		Algo: HS256.String(),
	}
)

// Protocol 协议字段
type Protocol struct {
	Type string `yaml:"t"`
	Algo string   `yaml:"a"`
}

// Constraint 约束字段
type Constraint struct {
	Signer      string `yaml:"sig"`
	Expiry      int64  `yaml:"exp"`
	Serial      string `yaml:"ser"`
	Beneficiary string `yaml:"ben"`
}

// YToken YAML Token处理体
type YToken struct {
	Protocol   *Protocol   `yaml:"p"`
	Constraint *Constraint `yaml:"c"`
}

// New 实例化一个 YToken
func New(constraint *Constraint) *YToken {
	return &YToken{
		Protocol: DefaultProtocol,
		Constraint: constraint,
	}
}

// Sign 进行签名
func (y *YToken) Sign(key []byte) (token string, err error) {
	yamlBytes, err := yaml.Marshal(y)
	if err != nil {
		return "", err
	}
	println(string(yamlBytes))

	mac := newAlgoType(y.Protocol.Algo).Func(key).Sign(yamlBytes)
	token = fmt.Sprintf("%s.%s", base64.RawURLEncoding.EncodeToString(yamlBytes), base64.RawURLEncoding.EncodeToString(mac))
	return
}

// Verify 验证签名，并将内容反序列化到当前结构体内
func (y *YToken) Verify(key, token []byte) (isEqual bool, err error) {
	stringList := strings.Split(*(*string)(unsafe.Pointer(&token)), ".")
	if len(stringList) != 2 {
		return false, errors.New("format of the token is wrong")
	}

	raw, err := base64.RawURLEncoding.DecodeString(stringList[0])
	if err != nil {
		return false, err
	}
	if err := yaml.Unmarshal(raw, y); err != nil {
		return false, err
	}

	mac, err := base64.RawURLEncoding.DecodeString(stringList[1])
	if err != nil {
		return false, err
	}

	return newAlgoType(y.Protocol.Algo).Func(key).Verify(raw, mac), nil
}

// Timeout 是否过期
func (y *YToken) Timeout() (isTimeout bool) {
	return y.Constraint.Expiry < time.Now().Unix()
}
