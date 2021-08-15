package gtls

import (
	"time"
)

const (
	KeyTypeRSA = KeyType(iota)
	KeyTypeEcdsa
	KeyTypeEd25519
)

type KeyType int8

// Options 选项配置
type Options struct {
	KeyType      KeyType
	Organization string
	Duration     time.Duration
	Hosts        []string

	Name string // 保存的文件名
}
