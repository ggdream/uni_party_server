package hashids

import "testing"

var hasher = New("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "abcd", 8)

func TestHashIDS_Encode(t *testing.T) {
	res := hasher.Encode(23166)
	println(res)
}

func TestHashIDS_Decode(t *testing.T) {
	res, err := hasher.Decode("3GdqE90k")
	if err != nil {
		panic(err)
	}
	println(res)
}
