package crypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

type rsaCrypto struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// RSA 实例化一个RSA操作器
func RSA(privateKeyPem []byte) (*rsaCrypto, error) {
	c := &rsaCrypto{}
	privateKey, err := c.UnmarshalPrivateKeyFromPem(privateKeyPem)
	if err != nil {
		return nil, err
	}

	c.privateKey = privateKey
	c.publicKey = &privateKey.PublicKey

	return c, nil
}

// Encrypt 加密明文
func (c *rsaCrypto) Encrypt(data []byte) (*wrapper, error) {
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, c.publicKey, data)
	if err != nil {
		return nil, err
	}
	return newWrapper(cipher), nil
}

// Decrypt 解密密文
func (c *rsaCrypto) Decrypt(data []byte, wrapperType wrapperType) ([]byte, error) {
	raw, err := getRawData(data, wrapperType)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, c.privateKey, raw)
}

// Sign 生成签名
func (c *rsaCrypto) Sign(data []byte) (*wrapper, error) {
	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		return nil, err
	}

	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return nil, err
	}
	return newWrapper(result), nil
}

// Verify 验证签名
func (c *rsaCrypto) Verify(data, sign []byte, wrapperType wrapperType) bool {
	var err error
	if sign, err = getRawData(sign, wrapperType); err != nil {
		return false
	}

	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		return false
	}

	return rsa.VerifyPKCS1v15(c.publicKey, crypto.SHA256, hash.Sum(nil), sign) == nil
}

// MarshalPrivateKeyToPem 将私钥序列化为Pem格式的字节切片
func (c *rsaCrypto) MarshalPrivateKeyToPem() (res bytes.Buffer, err error) {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.privateKey),
	}

	err = pem.Encode(&res, block)
	return
}

// MarshalPublicKeyToPem 将公钥序列化为Pem格式的字节切片
func (c *rsaCrypto) MarshalPublicKeyToPem() (res bytes.Buffer, err error) {
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(c.publicKey),
	}

	err = pem.Encode(&res, block)
	return
}

// UnmarshalPrivateKeyFromPem 从Pem文件反序列化出私钥
func (c *rsaCrypto) UnmarshalPrivateKeyFromPem(privateKeyPem []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPem)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// UnmarshalPublicKeyFromPem 从Pem文件反序列化出公钥
func (c *rsaCrypto) UnmarshalPublicKeyFromPem(publicKeyPem []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyPem)
	return x509.ParsePKCS1PublicKey(block.Bytes)
}
