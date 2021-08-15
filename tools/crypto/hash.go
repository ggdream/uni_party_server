package crypto

import (
	"hash"
	"io"
)

type hashX struct {
	hasher hash.Hash
}

// Hash 创建一个Hash操作器
func Hash(hasher hash.Hash) *hashX {
	return &hashX{
		hasher: hasher,
	}
}

// Data 对数据哈希，禁止并发调用
func (h *hashX) Data(data []byte) (*wrapper, error) {
	h.hasher.Reset()
	_, err := h.hasher.Write(data)
	if err != nil {
		return nil, err
	}
	return newWrapper(h.hasher.Sum(nil)), nil
}

// File 对文件哈希，禁止并发调用
func (h *hashX) File(file io.Reader) (*wrapper, error) {
	h.hasher.Reset()
	_, err := io.Copy(h.hasher, file)
	if err != nil {
		return nil, err
	}
	return newWrapper(h.hasher.Sum(nil)), nil
}
