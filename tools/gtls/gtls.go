package gtls

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"os"
	"time"
)


// New 创建gtls实例
func New(option *Options) (*Cert, error) {
	c := new(Cert)
	caCert, caKey, err := c.genCertificate(option.KeyType, option.Duration, option.Organization, option.Hosts, true)
	if err != nil {
		return nil, err
	}

	c.caCert = caCert
	c.caPrivateKey = caKey
	return c, err
}

// Default 用提供的CA证书创建gtls实例
func Default(caCert *x509.Certificate, caPrivateKey interface{}) *Cert {
	return &Cert{
		caCert:       caCert,
		caPrivateKey: caPrivateKey,
	}
}

// File 用CA文件创建gtls实例
func File(certFile, keyFile string) (*Cert, error) {
	certData, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	keyData, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	crtBlock, _ := pem.Decode(certData)
	keyBlock, _ := pem.Decode(keyData)

	crt, err := x509.ParseCertificate(crtBlock.Bytes)
	if err != nil {
		return nil, err
	}
	key, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return &Cert{
		caCert:       crt,
		caPrivateKey: key,
	}, nil
}

type Cert struct {
	// CA的证书
	caCert *x509.Certificate
	// CA的私钥
	caPrivateKey interface{}
}

// CA 获取CA的证书(.crt)和私钥(.key)的wrapper
func (c *Cert) CA() (*wrapper, *wrapper, error) {
	derBytes, err := x509.CreateCertificate(rand.Reader, c.caCert, c.caCert, c.getPublicKey(c.caPrivateKey), c.caPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(c.caPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	return newWrapper(derBytes, false), newWrapper(keyBytes, true), nil
}

// GenerateCrtAndKey 生成证书(.crt)和私钥(.key)的wrapper
func (c *Cert) GenerateCrtAndKey(option *Options) (*wrapper, *wrapper, error) {
	if option.KeyType == 0 {
		option.KeyType = KeyTypeRSA
	}
	if option.Duration == 0 {
		option.Duration = 365 * 24 * time.Hour
	}
	if option.Hosts == nil {
		option.Hosts = []string{"127.0.0.1"}
	}

	cert, privateKey, err := c.genCertificate(option.KeyType, option.Duration, option.Organization, option.Hosts, false)
	if err != nil {
		return nil, nil, err
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, cert, c.caCert, c.getPublicKey(privateKey), c.caPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}

	return newWrapper(derBytes, false), newWrapper(keyBytes, true), nil
}

// genCertificate 生成证书
func (c *Cert) genCertificate(keyType KeyType, duration time.Duration, organization string, hosts []string, isCA bool) (*x509.Certificate, interface{}, error) {
	serialNumber, err := c.genSerialNumber()
	if err != nil {
		return nil, nil, err
	}
	privateKey, err := c.genPrivateKey(keyType)
	if err != nil {
		return nil, nil, err
	}
	ip, dns := c.parseHosts(hosts)

	return &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{organization},
		},

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(duration),

		IPAddresses: ip,
		DNSNames:    dns,

		KeyUsage:              c.getKeyUsage(privateKey, isCA),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		IsCA:                  isCA,
		BasicConstraintsValid: true,
	}, privateKey, nil
}

// genSerialNumber 随机生成序列号
func (c *Cert) genSerialNumber() (*big.Int, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	return rand.Int(rand.Reader, serialNumberLimit)
}

// genPrivateKey 根据指定秘钥类型生成私钥
func (c *Cert) genPrivateKey(keyType KeyType) (key interface{}, err error) {
	switch keyType {
	case KeyTypeRSA:
		key, err = rsa.GenerateKey(rand.Reader, 2048)
	case KeyTypeEcdsa:
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case KeyTypeEd25519:
		_, key, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = errors.New("cannot find the key type")
	}
	return
}

// getPublicKey 从私钥中获取公钥
func (c *Cert) getPublicKey(privateKey interface{}) interface{} {
	switch k := privateKey.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

// getKeyUsage 设置秘钥用途
func (c *Cert) getKeyUsage(privateKey interface{}, isCA bool) x509.KeyUsage {
	keyUsage := x509.KeyUsageDigitalSignature

	if _, isRSA := privateKey.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}
	if isCA {
		keyUsage |= x509.KeyUsageCertSign
	}

	return keyUsage
}

// parseHosts 解析主机地址
func (c *Cert) parseHosts(hosts []string) (IPAddresses []net.IP, DNSNames []string) {
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			IPAddresses = append(IPAddresses, ip)
		} else {
			DNSNames = append(DNSNames, h)
		}
	}
	return
}
