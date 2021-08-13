package random

import (
	"crypto/md5"
	"fmt"
	"time"
)

// NewDeviceID 生成32位设备码
func NewDeviceID() string {
	data := fmt.Sprintf("%d:%s", time.Now().UnixNano(), Default())
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
