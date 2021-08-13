package safety

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

type ScryptAuth struct {
	Password	string
	Salt		string
}

func (s *ScryptAuth) Hash() (string, error) {
	res, err := scrypt.Key([]byte(s.Password), []byte(s.Salt), 1 << 15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(res), nil
}

func (s *ScryptAuth) Compare(hash string) (bool, error) {
	res, err := s.Hash()
	if err != nil {
		return false, err
	}
	return res == hash, nil
}
