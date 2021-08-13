package ytoken

import (
	"crypto/hmac"
	"crypto/sha256"
)

const (
	HS256 = AlgoType("HS256") // HMAC SHA256
)

// AlgoType 签名算法枚举
type AlgoType string

func newAlgoType(algo string) AlgoType {
	return AlgoType(algo)
}

// String 枚举转字符串
func (t AlgoType) String() string {
	switch t {
	case HS256:
		return "HS256"
	default:
		return "HS256"
	}
}

// Func 获取对应的哈希函数
func (t AlgoType) Func(key []byte) IAlgoFunc {
	switch t {
	case HS256:
		return &hs256{
			key: key,
		}
	default:
		return &hs256{
			key: key,
		}
	}
}

type IAlgoFunc interface {
	// Sign 进行签名
	Sign(raw []byte) (mac []byte)
	// Verify 验证签名
	Verify(raw, mac []byte) (isEqual bool)
}

type hs256 struct {
	key []byte
}

func (h *hs256) Sign(raw []byte) []byte {
	mac := hmac.New(sha256.New, h.key)
	mac.Write(raw)
	return mac.Sum(nil)
}

func (h *hs256) Verify(raw, mac []byte) bool {
	macer := hmac.New(sha256.New, h.key)
	macer.Write([]byte(raw))
	expectedMAC := macer.Sum(nil)

	return hmac.Equal(mac, expectedMAC)
}
