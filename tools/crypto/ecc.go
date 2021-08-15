package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"math/big"
)

type eccCrypto struct {
	eciesPrivateKey *ecies.PrivateKey
	eciesPublicKey  *ecies.PublicKey

	ecdsaPrivateKey *ecdsa.PrivateKey
	ecdsaPublicKey  *ecdsa.PublicKey
}

// ECC 实例化一个ECC操作器
func ECC(privateKeyPem []byte) (*eccCrypto, error) {
	c := &eccCrypto{}

	ecdsaPrivateKey, err := c.UnmarshalPrivateKeyFromPem(privateKeyPem)
	if err != nil {
		return nil, err
	}
	ecdsaPublicKey := &ecdsaPrivateKey.PublicKey

	eciesPublicKey := &ecies.PublicKey{
		X:      ecdsaPublicKey.X,
		Y:      ecdsaPublicKey.Y,
		Curve:  ecdsaPublicKey.Curve,
		Params: ecies.ParamsFromCurve(ecdsaPublicKey.Curve),
	}
	eciesPrivateKey := &ecies.PrivateKey{
		PublicKey: *eciesPublicKey,
		D:         ecdsaPrivateKey.D,
	}

	c.eciesPrivateKey = eciesPrivateKey
	c.eciesPublicKey = eciesPublicKey
	c.ecdsaPrivateKey = ecdsaPrivateKey
	c.ecdsaPublicKey = ecdsaPublicKey

	return c, nil
}

// Encrypt 加密明文
func (c *eccCrypto) Encrypt(data []byte) (*wrapper, error) {
	cipher, err := ecies.Encrypt(rand.Reader, c.eciesPublicKey, data, nil, nil)
	if err != nil {
		return nil, err
	}

	return newWrapper(cipher), nil
}

// Decrypt 解密密文
func (c *eccCrypto) Decrypt(data []byte, wrapperType wrapperType) ([]byte, error) {
	var err error
	if data, err = getRawData(data, wrapperType); err != nil {
		return nil, err
	}

	return c.eciesPrivateKey.Decrypt(data, nil, nil)
}

// Sign 生成签名
func (c *eccCrypto) Sign(data []byte) (*wrapper, *wrapper, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return nil, nil, err
	}

	r, s, err := ecdsa.Sign(rand.Reader, c.ecdsaPrivateKey, hash.Sum(nil))
	if err != nil {
		return nil, nil, err
	}
	rBytes, err := r.MarshalText()
	if err != nil {
		return nil, nil, err
	}
	sBytes, err := s.MarshalText()
	if err != nil {
		return nil, nil, err
	}

	return newWrapper(rBytes), newWrapper(sBytes), nil
}

// Verify 验证签名
func (c *eccCrypto) Verify(data, rBytes, sBytes []byte, wrapperType wrapperType) bool {
	var err error
	if rBytes, err = getRawData(rBytes, wrapperType); err != nil {
		return false
	}
	if sBytes, err = getRawData(sBytes, wrapperType); err != nil {
		return false
	}

	var r, s big.Int
	if err = r.UnmarshalText(rBytes); err != nil {
		return false
	}
	if err = s.UnmarshalText(sBytes); err != nil {
		return false
	}

	hash := sha256.New()
	hash.Write(data)

	return ecdsa.Verify(c.ecdsaPublicKey, hash.Sum(nil), &r, &s)
}

// GenerateKey 生成以太坊的公私钥和官方库的公私钥
func (c *eccCrypto) GenerateKey() (*ecies.PrivateKey, *ecies.PublicKey, *ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	eciesPrivateKey, err := ecies.GenerateKey(rand.Reader, elliptic.P256(), nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	eciesPublicKey := &eciesPrivateKey.PublicKey

	ecdsaPublicKey := &ecdsa.PublicKey{
		Curve: eciesPublicKey.Curve,
		X:     eciesPublicKey.X,
		Y:     eciesPublicKey.Y,
	}
	ecdsaPrivateKey := &ecdsa.PrivateKey{
		PublicKey: *ecdsaPublicKey,
		D:         eciesPrivateKey.D,
	}

	return eciesPrivateKey, eciesPublicKey, ecdsaPrivateKey, ecdsaPublicKey, nil
}

// MarshalPrivateKeyToPem 将私钥序列化为Pem格式的字节切片
func (c *eccCrypto) MarshalPrivateKeyToPem() (res bytes.Buffer, err error) {
	stream, err := x509.MarshalECPrivateKey(c.ecdsaPrivateKey)
	if err != nil {
		return
	}

	block := &pem.Block{
		Type:  "ECC PRIVATE KEY",
		Bytes: stream,
	}

	err = pem.Encode(&res, block)
	return
}

// MarshalPublicKeyToPem 将公钥序列化为Pem格式的字节切片
func (c *eccCrypto) MarshalPublicKeyToPem() (res bytes.Buffer, err error) {
	stream, err := x509.MarshalPKIXPublicKey(c.ecdsaPublicKey)
	if err != nil {
		return
	}

	block := &pem.Block{
		Type:  "ECC PUBLIC KEY",
		Bytes: stream,
	}

	err = pem.Encode(&res, block)
	return
}

// UnmarshalPrivateKeyFromPem 从Pem文件反序列化出私钥
func (c *eccCrypto) UnmarshalPrivateKeyFromPem(privateKeyPem []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPem)
	return x509.ParseECPrivateKey(block.Bytes)
}

// UnmarshalPublicKeyFromPem 从Pem文件反序列化出公钥
func (c *eccCrypto) UnmarshalPublicKeyFromPem(publicKeyPem []byte) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyPem)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*ecdsa.PublicKey), nil
}
