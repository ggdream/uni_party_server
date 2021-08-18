package crypto

// handler 业务处理器
type handler struct {
	rsa *rsaCrypto
	key, iv []byte
}
// NewHandler 实例化业务处理器
func NewHandler(privateKeyPem, key, iv []byte) (*handler, error) {
	rsa, err := RSA(privateKeyPem)
	if err != nil {
		return nil, err
	}

	return &handler{
		rsa: rsa,
		key: key,
		iv:  iv,
	}, nil
}

// Encrypt 加密数据
func (h *handler) Encrypt(data, randomKey []byte) (string, string, error) {
	encKey, err := h.rsa.Encrypt(randomKey)
	if err != nil {
		return "", "", err
	}
	
	aes1, err := AES(h.key, h.iv)
	if err != nil {
		return "", "", err
	}
	encTextTemp := aes1.Encrypt(data)
	aes2, err := AES(randomKey, h.iv)
	if err != nil {
		return "", "", err
	}
	encText := aes2.Encrypt(encTextTemp.Bytes())
	
	return encKey.ToBase64(), encText.ToBase64(), nil
}

// Decrypt 解密数据
func (h *handler) Decrypt(encKey, encText []byte) ([]byte, error) {
	key, err := h.rsa.Decrypt(encKey, WrapperBase64)
	if err != nil {
		return nil, err
	}

	aes1, err := AES(key, h.iv)
	if err != nil {
		return nil, err
	}
	dataTemp, err := aes1.Decrypt(encText, WrapperBase64)
	if err != nil {
		return nil, err
	}
	aes2, err := AES(h.key, h.iv)

	return aes2.Decrypt(dataTemp, WrapperNone)
}
