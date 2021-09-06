package pan

import "github.com/gin-gonic/gin"

// ShareController 分享文件/文件夹
func ShareController(c *gin.Context) {
	// TODO: 记录到Redis，并返回链接和秘钥
}

// GetShareController 获取别人的分享
func GetShareController(c *gin.Context) {
	// TODO: 验证链接和秘钥，返回临时Token和根目录
	// Token用户子目录访问和文件下载
}

// SaveShareController 保存别人的分享
func SaveShareController(c *gin.Context) {
	// TODO: 转存内容到自己的文件处，拷贝MongoDB的目录信息与自己一致
}
