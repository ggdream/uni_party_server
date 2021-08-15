package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type aesCrypto struct {
	key, iv []byte
	block cipher.Block
}

// AES 实例化一个AES-CBC操作器
func AES(key, iv []byte) (*aesCrypto, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &aesCrypto{
		key: key,
		iv: iv,
		block: block,
	}, nil
}

// Encrypt 加密明文
func (c *aesCrypto) Encrypt(data []byte) *wrapper {
	encrypter := cipher.NewCBCEncrypter(c.block, c.iv)
	data = c.PKCS5Padding(data, encrypter.BlockSize())
	result := make([]byte, len(data))
	encrypter.CryptBlocks(result, data)
	return newWrapper(result)
}

// Decrypt 解密密文
func (c *aesCrypto) Decrypt(data []byte, wrapperType wrapperType) ([]byte, error) {
	var err error
	if data, err = getRawData(data, wrapperType); err != nil {
		return nil, err
	}

	decrypter := cipher.NewCBCDecrypter(c.block, c.iv)
	result := make([]byte, len(data))
	decrypter.CryptBlocks(result, data)
	return c.PKCS5UnPadding(result), nil
}

// ZeroPadding 添加0填充
func (c *aesCrypto) ZeroPadding(data []byte, blockSize int) []byte {
	padNum := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{0}, padNum)...)
}

// ZeroUnPadding 去除0填充
func (c *aesCrypto) ZeroUnPadding(data []byte) []byte {
	return bytes.TrimFunc(data, func(r rune) bool { return r == rune(0) })
}

// PKCS5Padding 添加对齐填充
func (c *aesCrypto) PKCS5Padding(data []byte, blockSize int) []byte {
	padNum := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padNum)}, padNum)...)
}

// PKCS5UnPadding 去除对齐填充
func (c *aesCrypto) PKCS5UnPadding(data []byte) []byte {
	unNum := int(data[len(data)-1])
	return data[:(len(data) - unNum)]
}

// PKCS7Padding 添加对齐填充
func (c *aesCrypto) PKCS7Padding(data []byte, blockSize int) []byte {
	return c.PKCS5Padding(data, blockSize)
}

// PKCS7UnPadding 去除对齐填充
func (c *aesCrypto) PKCS7UnPadding(data []byte) []byte {
	return c.PKCS5UnPadding(data)
}
