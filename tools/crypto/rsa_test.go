package crypto

import (
	"fmt"
	"testing"
)

var c2, _ = RSA([]byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`))

func TestRsaCrypto_Encrypt(t *testing.T) {
	res, err := c2.Encrypt([]byte("我喜欢你"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ToBase64())
}

func TestRsaCrypto_Decrypt(t *testing.T) {
	res, err := c2.Decrypt([]byte("QBQ_0qBUifBRiePT80mH30_mJ7DZFxZkO58PvjpWei12I_GrDsizQRkRuhfCD2cQrXGOqQ16PYw2-EigatmXPRos-RakCrLIevuZB4uNoKBti1cxk-SWvt1STAvLhtYIuqIQfUXjMsXhACQGbFiDialcPFTMzmSR2WgZxpCMIdE"), WrapperBase64)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func TestRsaCrypto_Sign(t *testing.T) {
	res, err := c2.Sign([]byte("我喜欢你"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ToBase64())
}

func TestRsaCrypto_Verify(t *testing.T) {
	isEqual := c2.Verify([]byte("我喜欢你"), []byte("BjykeavwuD_8Zvw-STv4RKTJ7bsOAy-x9Brly9sBvto3DdfvEkSBpzRMkTmrrcf2WJGwIaJHo94tLxGT58EsAfyn5AhFlr_RmBzVpE2YrqGCPmhiaAIX7PJLoiaevMxBWO2aLxQx1fGuz54DH0FiIumx2RjRhouFC7TSMPJiVHw"), WrapperBase64)
	fmt.Println(isEqual)
}

func BenchmarkRsaCrypto_Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := c2.Encrypt([]byte("我喜欢你"))
		if err != nil {
			return
		}
	}
}

func BenchmarkRsaCrypto_Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := c2.Decrypt([]byte("QBQ_0qBUifBRiePT80mH30_mJ7DZFxZkO58PvjpWei12I_GrDsizQRkRuhfCD2cQrXGOqQ16PYw2-EigatmXPRos-RakCrLIevuZB4uNoKBti1cxk-SWvt1STAvLhtYIuqIQfUXjMsXhACQGbFiDialcPFTMzmSR2WgZxpCMIdE"), WrapperBase64)
		if err != nil {
			return
		}
	}
}

func BenchmarkRsaCrypto_Sign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := c2.Sign([]byte("我喜欢你"))
		if err != nil {
			return
		}
	}
}

func BenchmarkRsaCrypto_Verify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c2.Verify([]byte("我喜欢你"), []byte("BjykeavwuD_8Zvw-STv4RKTJ7bsOAy-x9Brly9sBvto3DdfvEkSBpzRMkTmrrcf2WJGwIaJHo94tLxGT58EsAfyn5AhFlr_RmBzVpE2YrqGCPmhiaAIX7PJLoiaevMxBWO2aLxQx1fGuz54DH0FiIumx2RjRhouFC7TSMPJiVHw"), WrapperBase64)
	}
}
