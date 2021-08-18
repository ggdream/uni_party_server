package middleware

const (
	KeyUID = "uid"	// 存储用户uid
	KeyDecValue = "decValue"	// 已使用crypto.Decrypt解密出的原始数据
	KeyEncValue = "encValue"	// 将使用crypto.Encrypt加密出的原始数据
)
