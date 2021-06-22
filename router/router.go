package router

import "github.com/gin-gonic/gin"

// Router 路由器
type Router struct {
	engine *gin.Engine
}

// New 创建一个路由器实例
func New(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (r *Router) Register() {
	externalRoutes := r.engine.Group("/x")
	erv1 := externalRoutes.Group("/v1")

	// v1 归档业务
	{
		archive := erv1.Group("/archive")

		// 点赞
		archive.POST("/star")

		// 添加评论
		archive.POST("/reply/add")

		// 取消评论
		archive.POST("/reply/del")

		// 收藏
		archive.POST("/collect")

		// 收集用户搜索行为
		archive.POST("/behavior/search")

		// 收集用户观看行为
		archive.POST("/behavior/time")
	}

	// v1 认证业务
	{
		auth := erv1.Group("/auth")

		// 密码登录
		auth.POST("/login/cipher")

		// 验证码登录
		auth.POST("/login/code")

		// 获取手机验证码
		auth.GET("/login/code/phone")

		// 获取邮箱验证码
		auth.GET("/login/code/email")

		// 学生认证
		auth.POST("/verify/student")

		// 公司认证
		auth.POST("/verify/company")

		// 高校认证
		auth.POST("/verify/university")

		// 机构认证
		auth.POST("/verify/institution")

		// 环境审核
		auth.GET("/safety/environment")
	}

	// v1 聊天业务
	{
		chat := erv1.Group("/chat")

		// 建立WebSocket连接
		chat.GET("/connect")
	}

	// v1 消息业务
	{
		event := erv1.Group("/events")

		// 获取订阅消息
		event.GET("/sub")

		// 删除消息
		event.POST("/delete")

		// 获取通知消息详情
		event.GET("/notice/detail")

		// 发布通知消息
		event.POST("/notice/publish")

		// 修改通知消息
		event.POST("/notice/repair")

		// 获取投票详情
		event.GET("/vote/detail")

		// 发布投票消息
		event.POST("/vote/publish")

		// 修改投票消息
		event.POST("/vote/repair")

		// 获取随机详情
		event.GET("/sortition/detail")

		// 发布随机消息
		event.POST("/sortition/publish")

		// 修改随机消息
		event.POST("/sortition/repair")

		// 获取报名详情
		event.GET("/participation/detail")

		// 发布报名消息
		event.POST("/participation/publish")

		// 修改报名消息
		event.POST("/participation/repair")

		// 搜索消息
		event.GET("/search")

		// 获取用户发布的消息
		event.GET("/users/publications")

		// 获取用户收藏的消息
		event.GET("/users/collections")

		// 点赞消息
		event.POST("/archive/star")

		// 给消息添加评论
		event.POST("/archive/reply/add")

		// 给消息删除评论
		event.POST("/archive/reply/del")

		// 收藏消息
		event.POST("/archive/collect")
	}

	// v1 辅助业务
	{
		helper := erv1.Group("/helpers")

		// 获取用户协议
		helper.GET("/protocols/user")

		// 获取服务协议
		helper.GET("/protocols/service")

		// 获取最新版本信息
		helper.GET("/version")

		// 获取所有版本信息
		helper.GET("/versions")

		// 关于我们
		helper.GET("/about/us")

		// 联系我们
		helper.GET("/contact")

		// 用户反馈
		helper.POST("/feedback")
	}

	// v1 用户业务
	{
		user := erv1.Group("/users")

		// 获取用户信息
		user.GET("/")

		// 获取用户的粉丝信息
		user.GET("/followers")

		// 获取用户的关注信息
		user.GET("/following")

		// 关注某个用户
		user.POST("/following/act")

		// 获取用户发布的消息
		user.GET("/events/publications")

		// 获取用户收藏的消息
		user.GET("/events/collections")

		// 获取用户发布的视频
		user.GET("/videos/publications")

		// 获取用户收藏的视频
		user.GET("/videos/collections")

		// 搜索
		user.GET("/search")

		// 更改用户头像
		user.POST("/profile/avatar")

		// 更改用户信息
		user.POST("/profile/update")

		// 获取主题色
		user.GET("/setttings/theme/get")

		// 设置主题色
		user.POST("/setttings/theme/set")

		// 获取邮箱推送状态
		user.GET("/setttings/push")

		// 设置消息推送
		user.POST("/setttings/push/event")

		// 设置视频推送
		user.POST("/setttings/push/video")
	}

	// v1 短视频业务
	{
		video := erv1.Group("/videos")

		// 获取后台推送的视频
		video.GET("/get")

		// 上传视频
		video.POST("/upload")

		// 删除视频
		video.POST("/delete")

		// 搜索视频
		video.GET("/search")

		// 获取用户发布的视频
		video.GET("/users/publications")

		// 获取用户收藏的视频
		video.GET("/users/collections")
	}
}
