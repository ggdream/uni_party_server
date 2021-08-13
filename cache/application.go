package cache

// Application 软件应用信息
type Application struct {}

func (a Application) joinAppUpdateKey() string {
	return "app:update"
}

// SetInfo 设置应用的最新信息
func (a Application) SetInfo(version, description string) error {
	return client.HMSet(a.joinAppUpdateKey(), "version", version, "description", description)
}

// GetInfo 获取应用的最新信息
func (a Application) GetInfo() (interface{}, error) {
	return client.HMGet(a.joinAppUpdateKey())
}
