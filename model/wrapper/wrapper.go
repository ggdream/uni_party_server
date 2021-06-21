package base


// WrapperModel 返回模型的外部包装体
type WrapperModel struct {
	Code    string      `json:"code" form:"code"`
	Data    interface{} `json:"data" form:"data"`
	TTL     int         `json:"ttl" form:"ttl"`
	Message string      `json:"message" form:"message"`
}
