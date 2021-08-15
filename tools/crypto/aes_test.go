package crypto

import (
	"fmt"
	"testing"
)

var c1, _ = AES([]byte("0102030405060708"), []byte("0102030405060708"))


func TestAesCrypto_Encrypt(t *testing.T) {
	data := c1.Encrypt([]byte("我喜欢你"))
	fmt.Println(data.ToBase64())
}

func TestAesCrypto_Decrypt(t *testing.T) {
	data, err := c1.Decrypt([]byte("l8JfKWDxd4GdGyDlpuLnSg"), WrapperBase64)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func BenchmarkAesCrypto_Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c1.Encrypt([]byte("我喜欢你"))
	}
}

func BenchmarkAesCrypto_Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := c1.Decrypt([]byte("l8JfKWDxd4GdGyDlpuLnSg"), WrapperBase64)
		if err != nil {
			return
		}
	}
}
