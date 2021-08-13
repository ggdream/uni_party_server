package cache

// Contact 联系我们的方式
type Contact struct {}

func (c Contact) joinInfoKey() string {
	return "we:contact"
}

// SetInfo 设置应用的最新信息
func (c Contact) SetInfo(version, description string) error {
	return client.HMSet(c.joinInfoKey(), "version", version, "description", description)
}

// GetInfo 获取应用的最新信息
func (c Contact) GetInfo() (interface{}, error) {
	return client.HMGet(c.joinInfoKey())
}
