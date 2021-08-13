package ytoken

// Handler YToken 处理器
type Handler struct {
	key []byte
}

// NewHandler 实例化一个 YToken 处理器
func NewHandler(key []byte) *Handler {
	return &Handler{
		key: key,
	}
}

func (h *Handler) Sign(yToken *YToken) (token string, err error) {
	return yToken.Sign(h.key)
}

func (h *Handler) Verify(token string) (yToken *YToken, isEqual bool, err error) {
	temp := &YToken{}
	isEqual, err = temp.Verify(h.key, []byte(token))
	yToken = temp
	return
}
