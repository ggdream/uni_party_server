package sms

import "testing"

func TestMail_Send(t *testing.T) {
	mailer := NewMailer("smtp.126.com", 465, "mocaraka@126.com", "YMIRUQKQUVHCVDOB")
	if err := mailer.Post("这是你的登录验证码，有效期限为五分钟，请勿泄露给他人", "<h1>你在干嘛呢~~</h1>", "1586616064@qq.com"); err != nil {
		panic(err)
	}
}
