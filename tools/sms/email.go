package sms

import (
	gomail "gopkg.in/mail.v2"
)

// Mail 邮箱消息发送体
type Mail struct {
	mail   string
	dialer *gomail.Dialer
}

// NewMailer 实例化一名邮递员
func NewMailer(host string, port int, mail string, password string) *Mail {
	return &Mail{
		mail:   mail,
		dialer: gomail.NewDialer(host, port, mail, password),
	}
}

// Post 给一个或着多个邮箱发送邮件
func (m *Mail) Post(subject, content string, receiver ...string) error {
	return m.dialer.DialAndSend(m.pack(subject, content, receiver...))
}

// pack 为邮件添加各项信息
func (m *Mail) pack(subject, content string, receiver ...string) *gomail.Message {
	n := gomail.NewMessage()
	n.SetHeaders(map[string][]string{
		"From":    {m.mail},
		"To":      receiver,
		"Subject": {subject},
	})
	n.SetBody("text/html", content)
	return n
}
