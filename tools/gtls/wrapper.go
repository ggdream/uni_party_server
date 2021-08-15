package gtls

import (
	"bytes"
	"encoding/pem"
	"os"
)

type wrapper struct {
	data  []byte
	isKey bool // key or crt
}

func newWrapper(data []byte, isKey bool) *wrapper {
	return &wrapper{
		data:  data,
		isKey: isKey,
	}
}

// Raw 获取原始证书字节切片
func (w *wrapper) Raw() []byte {
	return w.data
}

// Pem 将证书序列化为Pem格式
func (w *wrapper) Pem() (data bytes.Buffer, err error) {
	stream := &pem.Block{
		Type:  w.getType(),
		Bytes: w.data,
	}

	err = pem.Encode(&data, stream)
	return
}

// File 将内容写入文件
func (w *wrapper) File(name string) error {
	data, err := w.Pem()
	if err != nil {
		return err
	}

	return os.WriteFile(name, data.Bytes(), 0644)
}

// getType 获取Type字段的值
func (w *wrapper) getType() string {
	if w.isKey {
		return "PRIVATE KEY"
	} else {
		return "CERTIFICATE"
	}
}
