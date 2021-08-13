package hashids

import "github.com/speps/go-hashids/v2"

// HashIDS go-hashids包的封装
type HashIDS struct {
	hd	*hashids.HashID
}

// New 实例化一个 HashIDS
func New(alphabet, salt string, minLength int) *HashIDS {
	data := &hashids.HashIDData{
		Alphabet: alphabet,
		MinLength: minLength,
		Salt: salt,
	}
	hd, err := hashids.NewWithData(data)
	if err != nil {
		panic(err)
	}
	return &HashIDS{
		hd: hd,
	}
}

// Encode 编码
func (i *HashIDS) Encode(data uint) (hash string) {
	hash, _ = i.hd.EncodeInt64([]int64{int64(data)})
	return
}

// Decode 解码
func (i *HashIDS) Decode(hash string) (int64, error) {
	res, err := i.hd.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}
	return res[0], nil
}
