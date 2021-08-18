package crypto

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewHandler(t *testing.T) {
	h, err := NewHandler(
		[]byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCWBRyN1W2GpBBV8+K+WOQX0sCQI5tV2LZx43t3DcNd5nQYe6Oi
c9Rys7oJxTUDh2xXSpRI7tgw7S2nMJQiWGr2c7ul83lfXonPSs+6XmRWUeKFpk1l
s8izctcUVNJtXlntDuOtAXnNbSxPS7jqCfIn3/yDHvsG9TUw3J96f7EShQIDAQAB
AoGAANf1rnN/NUF6F0L6BXZiVb9nj01FSiZRyzXEWRuCZl8/Wig51EC5xOl2bUR4
tSaRKRdHTFp4ls4Q2aB3/+P/jxVnKKRAO+QioGsCztyjmzKksE1uHT9fm7eafh92
BybwhuSWnpoaQu/mplgeKy9dqHTNRNYhOab9bWSeftMPmmECQQDS6Rwfc+z0GUjf
ubGnukPGSpt/GXnZ0AmqgIgC/PJPjTUS8s4CrU0dSpoQam+7RmJHhbR+vjgtQOx1
R4jMixIhAkEAtheHepzmJcWkLHLW8oEL8kTiEBEfeg367fReALl7FJRzwCc7E2y1
Jbya5lFpoYB/rOpyzwuehtUveU7z6xR75QJANkzwAaUOB5aO7ZLYi29oTaAdwq5j
cdGe+3fWDLblB9g6JuO+aHAH35e6bKmKlLO9T967proAqp4Bgvxk2wyfIQJAA/4e
FV0tWrSrteHKvsR/m49RGaMWepml3+PUz9VBUqEyrIrvq6xdiHdLOjOcylLD+emm
HM4XasNv6AwIIyZu1QJAIb0VsJyw8PHRrMjsIP9lTf+BQctNdTLT9lj9yed+l7wx
ACv/ZM9xZAZXJdqAzYByq8YoStRwpeMOJ5/H6Fo4jg==
-----END RSA PRIVATE KEY-----`),
		[]byte("0102030405060708"),
		[]byte("0102030405060708"),
	)
	if err != nil {
		panic(err)
	}

	encKey, encText, err := h.Encrypt([]byte("我喜欢你"), []byte("0102030405060708"))
	if err != nil {
		panic(err)
	}
	fmt.Println(encKey)
	fmt.Println(encText)

	rawText, err := h.Decrypt([]byte(encKey), []byte(encText))
	if err != nil {
		panic(err)
	}

	fmt.Println(*(*string)(unsafe.Pointer(&rawText)))
}
