package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"unsafe"
)

const (
	WrapperNone = wrapperType(iota)
	WrapperHex
	WrapperBase64
)

type (
	wrapperType int8
	wrapper     struct {
		data []byte
	}
)

func newWrapper(data []byte) *wrapper {
	return &wrapper{
		data: data,
	}
}

// Bytes 获取原始数据字节切片
func (w *wrapper) Bytes() []byte {
	return w.data
}

// ToHex 原始数据转16进制字符串
func (w *wrapper) ToHex() string {
	return hex.EncodeToString(w.data)
}

// ToBase64 原始数据转Base64字符串
func (w *wrapper) ToBase64() string {
	return base64.RawURLEncoding.EncodeToString(w.data)
}

// FromHex 16进制字符串转原始数据
func (w *wrapper) FromHex() ([]byte, error) {
	return hex.DecodeString(*(*string)(unsafe.Pointer(&w.data)))
}

// FromBase64 Base64字符串转原始数据
func (w *wrapper) FromBase64() ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(*(*string)(unsafe.Pointer(&w.data)))
}

// getRawData 获取原始数据
func getRawData(data []byte, wrapperType wrapperType) (raw []byte, err error) {
	wrapper := newWrapper(data)
	switch wrapperType {
	case WrapperHex:
		raw, err = wrapper.FromHex()
	case WrapperBase64:
		raw, err = wrapper.FromBase64()
	default:
		raw = wrapper.Bytes()
	}

	return
}
