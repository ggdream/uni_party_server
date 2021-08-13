package random

import "math/rand"

const (
	number = "0123456789"
	numberLength = len(number)
)

// NewSMSCode 生成6位数字消息验证码
func NewSMSCode() (res string) {
	for i := 0; i < 6; i++ {
		res += string(number[rand.Intn(numberLength)])

	}
	return
}
