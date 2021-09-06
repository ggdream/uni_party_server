package oss

type Model struct {
	Name string
	Hash string
	Size int64
	Type string

	LastModified int64
	ContentType  string
}

type DelModel struct {
	IsBlob bool
	Prefix string
}
