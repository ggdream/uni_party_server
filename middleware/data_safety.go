package middleware

import (
	"gateway/tools/crypto"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

var (
	iv = []byte("0102030405060708")
)

// DataEncrypt 给响应数据解密
func DataEncrypt() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("enc_key")
		if key == "" || len(key) != 16 {
			errno.Abort(c, errno.TypeEncKeyGetFailed)
			return
		}
		aes, err := crypto.AES([]byte(key), iv)
		if err != nil {
			errno.Abort(c, errno.TypeCryptoInstanceFailed)
			return
		}

		c.Next()

		data := aes.Encrypt([]byte(c.GetString(KeyEncValue)))
		errno.Perfect(c, map[string]interface{}{"enc": data.ToBase64()})
	}
}

func DataDecrypt() gin.HandlerFunc {
	hander, err := crypto.NewHandler(
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

	return func(c *gin.Context) {
		var form struct {
			EncKey  string `json:"encKey" form:"enc_key" binding:"required"`
			EncText string `json:"encText" form:"enc_text" binding:"required"`
		}
		if err := c.ShouldBind(&form); err != nil {
			errno.Abort(c, errno.TypeParamsParsingErr)
			return
		}

		data, err := hander.Decrypt([]byte(form.EncKey), []byte(form.EncText))
		if err != nil {
			errno.Abort(c, errno.TypeDecryptFailed)
			return
		}

		c.Set(KeyDecValue, data)
		c.Next()
	}
}
