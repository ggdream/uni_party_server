package cache

import (
	"errors"
	"fmt"
	"time"
)

// SMS 消息验证码缓存与验证
type SMS struct{}

// joinDeviceAlreadyKey 获取保存在指定时间内已经发送了验证码的设备的string-key
func (s SMS) joinDeviceAlreadyKey(account string) string {
	return fmt.Sprintf("sms:send:%s", account)
}

// joinCodeKey 获取存储smsCode的string-key
func (s SMS) joinCodeKey(account string) string {
	return fmt.Sprintf("sms:code:%s", account)
}

// isAlreadySendCode 判断在指定期限内是否发过验证码
func (s SMS) isAlreadySendCode(account string) (bool, error) {
	return client.Exists(s.joinDeviceAlreadyKey(account))
}

// setDeviceAlready 设置key，保证指定时间内无法发送第二次验证码
func (s SMS) setDeviceAlready(account string) error {
	return client.Set(s.joinDeviceAlreadyKey(account), 0, time.Minute)
}

// joinCodeValue 获取code对应的value
func (s SMS) joinCodeValue(deviceCode, smsCode string) string {
	return fmt.Sprintf("%s:%s", deviceCode, smsCode)
}

// setCode 缓存验证码，有效期：5min
func (s SMS) setCode(account, deviceCode, smsCode string) error {
	return client.Set(s.joinCodeKey(account), s.joinCodeValue(deviceCode, smsCode), 5*time.Minute)
}

// Save 存储验证码
func (s SMS) Save(account, deviceCode, smsCode string) error {
	// 一分钟内是否发送过验证码
	isExist, err := s.isAlreadySendCode(account)
	if err != nil {
		return err
	} else if isExist {
		return errors.New("the device had send a code with expire")
	}

	// 若没有，则缓存
	if err := s.setDeviceAlready(account); err != nil {
		return err
	}
	// 如果距离上次发送时间不到5min，则直接覆盖原有code
	return s.setCode(deviceCode, smsCode, account)
}

// getCodeValue 从Redis中获取code对应的value
func (s SMS) getCodeValue(account string) (string, error) {
	return client.Get(s.joinCodeKey(account)).Result()
}

// delSaveCodeKey 验证成功后，删除保留的key
func (s SMS) delSaveCodeKey(account string) error {
	_, err := client.Del(s.joinCodeKey(account))
	return err
}

// Verify 验证设备码与验证码是否匹配
func (s SMS) Verify(account, deviceCode, smsCode string) (bool, error) {
	code, err := s.getCodeValue(account)
	if err != nil {
		return false, err
	}

	// 判断从Redis中获取到的与从客户端获取到的是否一致
	if code != s.joinCodeValue(deviceCode, smsCode) {
		return false, nil
	}

	_ = s.delSaveCodeKey(account)
	return true, nil
}
