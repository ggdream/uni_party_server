package pan

import "github.com/gin-gonic/gin"

// UploadSingleController 上传单个文件
func UploadSingleController(c *gin.Context) {
	// TODO: 解析出文件基本信息：名称，大小，类型
	// 生成全局唯一ID，作为文件名
	// io.Reader至OSS
	// 记录到MongoDB
}

// UploadMultiController 上传多个文件
func UploadMultiController(c *gin.Context) {
	// TODO: ....
}
