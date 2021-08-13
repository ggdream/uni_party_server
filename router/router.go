package router

import (
	chat2 "gateway/controller/chat"
	"github.com/gin-gonic/gin"
)

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
	archive := erv1.Group("/archive")
	{

		// 点赞
		archive.POST("/star")				// 视频点赞表、视频点赞列表、视频表

		// 添加评论
		archive.POST("/reply/add")		// 视频祖宗评论表、视频子孙评论表

		// 取消评论
		archive.POST("/reply/del")		// 视频祖宗评论表、视频子孙评论表

		// 收藏
		archive.POST("/collect")			// 视频收藏表

		// 收集用户搜索行为
		archive.POST("/behavior/search")

		// 收集用户观看行为
		archive.POST("/behavior/time")
	}

	// v1 认证业务
	auth := erv1.Group("/auth")
	{

		// 密码登录
		auth.POST("/login/cipher")		// 用户登录表

		// 验证码登录
		auth.POST("/login/code")			// 用户登录表[、用户地址表+用户信息表]

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
	chat := erv1.Group("/chat")
	{

		// 建立WebSocket连接
		chat.GET("/connect", chat2.ChatHandler)				// 聊天消息表(单聊) OR 聊天消息表(群聊)
	}

	// v1 消息业务
	event := erv1.Group("/events")
	{

		// 获取订阅消息
		event.GET("/sub")						// 用户订阅列表

		// 删除消息
		event.POST("/delete")					// 消息总表、用户订阅列表

		// 获取通知消息详情
		event.GET("/notice/detail")			// 用户订阅列表、用户发布列表、用户发布内容、消息总表

		// 发布通知消息
		event.POST("/notice/publish")			// 消息总表、用户发布列表

		// 修改通知消息
		event.POST("/notice/repair")			// 消息总表、用户发布内容表

		// 获取投票详情
		event.GET("/vote/detail")				// 用户订阅列表、用户发布列表、用户发布内容、消息总表

		// 发布投票消息
		event.POST("/vote/publish")			// 消息总表、用户发布列表

		// 修改投票消息
		event.POST("/vote/repair")			// 消息总表、用户发布内容表

		// 获取随机详情
		event.GET("/sortition/detail")		// 用户订阅列表、用户发布列表、用户发布内容、消息总表

		// 发布随机消息
		event.POST("/sortition/publish")		// 消息总表、用户发布列表

		// 修改随机消息
		event.POST("/sortition/repair")		// 消息总表、用户发布内容表

		// 获取报名详情
		event.GET("/participation/detail")	// 用户订阅列表、用户发布列表、用户发布内容、消息总表

		// 发布报名消息
		event.POST("/participation/publish")	// 消息总表、用户发布列表

		// 修改报名消息
		event.POST("/participation/repair")	// 消息总表、用户发布内容表

		// 搜索消息
		event.GET("/search")					// 用户订阅列表、ES匹配

		// 获取用户发布的消息
		event.GET("/users/publications")

		// 获取用户收藏的消息
		event.GET("/users/collections")

		// Get消息
		event.POST("/archive/star")			// 确认消息

		// Deprecated
		// 给消息添加评论
		event.POST("/archive/reply/add")

		// Deprecated
		// 给消息删除评论
		event.POST("/archive/reply/del")

		// 关注消息
		event.POST("/archive/collect")		// 用户的关注列表、消息的关注列表
	}

	// v1 辅助业务
	helper := erv1.Group("/helpers")
	{

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
	user := erv1.Group("/users")
	{

		// 获取用户信息
		user.GET("/")

		// 获取用户的粉丝信息
		user.GET("/followers")		// 关注表、用户信息表

		// 获取用户的关注信息
		user.GET("/following")		// 关注表、用户信息表

		// 关注某个用户
		user.POST("/following/act")	// 关注表、用户信息表

		// 获取好友（双向关注）
		user.GET("/friends")			// 关注表、用户信息表

		// 获取用户发布的消息
		user.GET("/events/publications")	// 用户的发布内容、消息总表

		// 获取用户关注的消息
		user.GET("/events/collections")	// 用户的关注列表、消息的关注列表

		// 获取用户发布的视频
		user.GET("/videos/publications")	// 用户发布视频缓存表、视频表

		// 获取用户收藏的视频
		user.GET("/videos/collections")	// 用户收藏视频缓存表、视频收藏表

		// 搜索
		user.GET("/search")				// 用户缓存表、用户信息表、ES模糊匹配

		// 更改用户头像
		user.POST("/profile/avatar")		// 用户信息表、用户缓存表

		// 更改用户信息
		user.POST("/profile/update")		// 用户信息表、用户缓存表

		// 获取主题色
		user.GET("/setttings/theme/get")		// 应用设置表

		// 设置主题色
		user.POST("/setttings/theme/set")		// 应用设置表

		// 获取邮箱推送状态
		user.GET("/setttings/push")			// 应用设置表

		// 设置消息推送
		user.POST("/setttings/push/event")	// 应用设置表

		// 设置视频推送
		user.POST("/setttings/push/video")	// 应用设置表
	}

	// v1 短视频业务
	video := erv1.Group("/videos")
	{

		// 获取后台推送的视频
		video.GET("/get")

		// 上传视频
		video.POST("/upload")				// 视频上传待审核表

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
