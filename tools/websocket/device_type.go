package websocket

const (
	// 网页端
	Web DeviceType = iota
	// 移动端
	Mobile
	// 桌面端
	Desktop
)

// DeviceType 用户设备类型枚举
type DeviceType int8

// NewDeviceType 将int8类型转为Enum类型
func NewDeviceType(index int8) DeviceType {
	return DeviceType(index)
}

// Index 将Enum类型转为int8类型
func (c DeviceType) Index() int8 {
	return int8(c)
}
