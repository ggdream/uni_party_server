package crypto

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

func TestHashX_Data(t *testing.T) {
	h := Hash(sha256.New())
	for i := 0; i < 20000; i++ {
		res, err := h.Data([]byte("我喜欢你"))
		if err != nil {
			panic(err)
		}
		fmt.Println(res.ToBase64())
	}
}

func BenchmarkHashX_File(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./hash_test.go")
		if err != nil {
			panic(err)
		}

		h := Hash(sha256.New())
		res, err := h.File(file)
		if err != nil {
			panic(err)
		}
		res.ToBase64()
	}
}
