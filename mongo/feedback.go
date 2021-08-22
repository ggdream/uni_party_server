package mongo

const (
	fdCollectionName = "feedback"
)

// FeedbackDocument 用户反馈元信息
type FeedbackDocument struct {
	UID   uint   `json:"uid"`
	Email string `json:"email"`

	Type    int8   `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Insert 插入新记录
func (f *FeedbackDocument) Insert() error {
	return client.Insert(fdCollectionName, f)
}
