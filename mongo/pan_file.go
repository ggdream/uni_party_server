package mongo

// PanFileDocument 用户持有文件
type PanFileDocument struct {
	UID        uint
	Name       string
	Hash       string
	ObjectName string
}
