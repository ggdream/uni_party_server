package crypto

import "testing"

var w = newWrapper([]byte("我喜欢你"))

func BenchmarkWrapper_FromBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w.ToBase64()
	}
}

func BenchmarkWrapper_ToHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w.ToHex()
	}
}
