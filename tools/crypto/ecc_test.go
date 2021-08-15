package crypto

import (
	"fmt"
	"testing"
)

var c3, _ = ECC([]byte(`-----BEGIN ECC PRIVATE KEY-----
MHcCAQEEIP5+QjDwv6jKx6YRpneGcZi25oJP2N1RnBQb0vzNBvd+oAoGCCqGSM49
AwEHoUQDQgAEx0HF+5iyI1S1f8lqC8P3ZXXVuLfs+AnFGzrhDhxY1ErYhlBdOWVV
zF60SjZz0eHxa6AU7AYnyMjMpWDWCOk7dQ==
-----END ECC PRIVATE KEY-----`))

func TestEccCrypto_GenerateKey(t *testing.T) {
	c := new(eccCrypto)
	_, _, key, _, err := c.GenerateKey()
	if err != nil {
		panic(err)
	}

	c.ecdsaPrivateKey = key
	pem, err := c.MarshalPrivateKeyToPem()
	if err != nil {
		panic(err)
	}

	fmt.Println(pem.String())
}

func TestEccCrypto_Encrypt(t *testing.T) {
	res, err := c3.Encrypt([]byte("我喜欢你"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ToBase64())
}

func TestEccCrypto_Decrypt(t *testing.T) {
	res, err := c3.Decrypt([]byte("BHercwtkFja2wSHtvWdnPjTE6A5yZjdkj8yUWy8VfuKHnnkRlrTNZbD8YOH_PuerG1mt39h8QDWAAdymRj_wMkMeOvv04ZKiyS16aYf0w6fIROxRLVXeMBAfW854QPoZ2xUTIigDpN6fnx7DzKFi9n7a84qbpHsPvUm4xiA"), WrapperBase64)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func TestEccCrypto_Sign(t *testing.T) {
	r, s, err := c3.Sign([]byte("我喜欢你"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("r:%s\n", r.ToBase64())
	fmt.Printf("s:%s\n", s.ToBase64())
}

func TestEccCrypto_Verify(t *testing.T) {
	isEqual := c3.Verify([]byte("我喜欢你"), []byte("NDY4NTE3NTY3Mzc0Njk1NzQ1MzgxNDE3NTA4ODg0MTY2ODM2OTI4MzAxOTcyNDc3ODAzNzQ0NDA2OTEyNDE4MzkxNDAxNTc3MDk0MQ"), []byte("NTgwNTE3MjA4MzAxNDk5NDc0NTM4ODA4MjkwODAwMDc5MDg4NjQ5ODgyMjY4ODMzNTgzMjc1ODE0ODk1NzczNzA1MDgzMzQ5NjE1ODQ"), WrapperBase64)
	fmt.Println(isEqual)
}

func BenchmarkEccCrypto_Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := c3.Encrypt([]byte("我喜欢你"))
		if err != nil {
			return
		}
	}
}

func BenchmarkEccCrypto_Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := c3.Decrypt([]byte("BHercwtkFja2wSHtvWdnPjTE6A5yZjdkj8yUWy8VfuKHnnkRlrTNZbD8YOH_PuerG1mt39h8QDWAAdymRj_wMkMeOvv04ZKiyS16aYf0w6fIROxRLVXeMBAfW854QPoZ2xUTIigDpN6fnx7DzKFi9n7a84qbpHsPvUm4xiA"), WrapperBase64)
		if err != nil {
			return
		}
	}
}

func BenchmarkEccCrypto_Sign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := c3.Sign([]byte("我喜欢你"))
		if err != nil {
			return
		}
	}
}

func BenchmarkEccCrypto_Verify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c3.Verify([]byte("我喜欢你"), []byte("NDY4NTE3NTY3Mzc0Njk1NzQ1MzgxNDE3NTA4ODg0MTY2ODM2OTI4MzAxOTcyNDc3ODAzNzQ0NDA2OTEyNDE4MzkxNDAxNTc3MDk0MQ"), []byte("NTgwNTE3MjA4MzAxNDk5NDc0NTM4ODA4MjkwODAwMDc5MDg4NjQ5ODgyMjY4ODMzNTgzMjc1ODE0ODk1NzczNzA1MDgzMzQ5NjE1ODQ"), WrapperBase64)
	}
}
