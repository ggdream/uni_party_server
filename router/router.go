package router

import (
	"gateway/controller/auth"
	"gateway/controller/chat"
	"gateway/controller/event"
	"gateway/controller/helper"
	"gateway/controller/user"
	"gateway/controller/video"

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

	// v1 视频归档业务
	archiveRouter := erv1.Group("/archive")
	{
		// 点赞
		archiveRouter.POST("/star")				// 视频点赞表、视频点赞列表、视频表

		// 收藏
		archiveRouter.POST("/collect")			// 视频收藏表


		// 添加评论
		archiveRouter.POST("/reply/add")		// 视频祖宗评论表、视频子孙评论表

		// 取消评论
		archiveRouter.POST("/reply/del")		// 视频祖宗评论表、视频子孙评论表

		// 获取评论
		archiveRouter.POST("/reply/get")		// 视频祖宗评论表、视频子孙评论表

		// 获取子评论
		archiveRouter.POST("/reply/son")		// 视频祖宗评论表、视频子孙评论表

		// // 我喜欢这个，多来点儿
		// archiveRouter.POST("/love/yes", )

		// // 我不喜欢这个，少来
		// archiveRouter.POST("/love/no", )


		// // 收集用户搜索行为
		// archiveRouter.POST("/behavior/search")

		// // 收集用户观看行为
		// archiveRouter.POST("/behavior/time")


		// 当前热点
		archiveRouter.GET("/hot")
	}

	// v1 认证业务
	authRouter := erv1.Group("/auth")
	{
		// 密码登录
		authRouter.POST("/login/cipher", auth.LoginByCipherController)		// 用户登录表

		// 验证码登录
		authRouter.POST("/login/code", auth.LoginByCodeController)			// 用户登录表[、用户地址表+用户信息表]

		// 获取手机验证码
		authRouter.GET("/login/code/phone", auth.CodePhoneController)

		// 获取邮箱验证码
		authRouter.GET("/login/code/email", auth.CodeEMailController)

		// 学生认证，移动端
		authRouter.POST("/verify/student")

		// 公司认证，网页端
		authRouter.POST("/verify/company")

		// 高校认证，网页端
		authRouter.POST("/verify/university")
	}

	// v1 聊天业务
	chatRouter := erv1.Group("/chat")
	{
		// 建立WebSocket连接
		chatRouter.GET("/connect", chat.ChatHandler)				// 聊天消息表(单聊) OR 聊天消息表(群聊)
	}

	// v1 消息业务
	eventRouter := erv1.Group("/event")
	{
		// 获取订阅消息
		eventRouter.GET("/sub", event.SubscribeController)						// 用户订阅列表

		// 删除消息
		eventRouter.POST("/delete", event.DeleteController)					// 消息总表、用户订阅列表

		// 获取消息详情
		eventRouter.GET("/detail", event.DetailController)			// 用户订阅列表、用户发布列表、用户发布内容、消息总表


		// 发布通知消息
		eventRouter.POST("/n/publish", event.NoticeCreateController)			// 消息总表、用户发布列表

		// 修改通知消息
		eventRouter.POST("/n/repair", event.NoticeUpdateController)			// 消息总表、用户发布内容表


		// 发布投票详情
		eventRouter.GET("/v/publish", event.VoteCreateController)				// 用户订阅列表、用户发布列表、用户发布内容、消息总表

		// 修改投票消息
		eventRouter.POST("/v/repair", event.VoteUpdateController)			// 消息总表、用户发布列表


		// 发布随机消息
		eventRouter.POST("/s/publish", event.SortitionCreateController)		// 消息总表、用户发布列表

		// 修改随机消息
		eventRouter.POST("/s/repair", event.SortitionUpdateController)		// 消息总表、用户发布内容表


		// 发布报名消息
		eventRouter.POST("/p/publish", event.ParticipationCreateController)	// 消息总表、用户发布列表

		// 修改报名消息
		eventRouter.POST("/p/repair", event.ParticipationUpdateController)	// 消息总表、用户发布内容表


		// 搜索消息
		eventRouter.GET("/search", event.SearchController)					// 用户订阅列表、ES匹配

		// 获取用户发布的消息
		eventRouter.GET("/users/public", event.PublicationGetController)

		// 获取用户关注的消息
		eventRouter.GET("/users/attend", event.AttendGetController)


		// Get消息
		eventRouter.POST("/archive/get", event.ArchiveGetController)

		// 关注消息
		eventRouter.POST("/archive/attend/add", event.ArchiveAddAttendController)

		// 取关消息
		eventRouter.POST("/archive/attend/del", event.ArchiveDelAttendController)
	}

	// v1 辅助业务
	helperRouter := erv1.Group("/helper")
	{
		// 获取应用版本信息
		helperRouter.GET("/app", helper.UpdateController)

		// 联系我们
		helperRouter.GET("/contact", helper.ContactController)

		// 用户反馈
		helperRouter.POST("/feedback", helper.FeedbackController)
	}

	// v1 用户业务
	userRouter := erv1.Group("/user")
	{
		// 获取用户信息
		userRouter.GET("/", user.GetProfileController)

		// 获取用户的粉丝信息
		userRouter.GET("/followers", user.FollowersController)		// 关注表、用户信息表

		// 获取用户的关注信息
		userRouter.GET("/following", user.FollowingController)		// 关注表、用户信息表

		// 关注某个用户
		userRouter.POST("/follow", user.FollowItController)	// 关注表、用户信息表

		// // 获取好友（双向关注）
		// userRouter.GET("/friends")			// 关注表、用户信息表

		// 搜索
		userRouter.GET("/search", user.SearchController)				// 用户缓存表、用户信息表、ES模糊匹配

		// 更改用户头像
		userRouter.POST("/profile/avatar", user.UploadAvatarController)		// 用户信息表、用户缓存表

		// 更改用户信息
		userRouter.POST("/profile/update", user.SetProfileController)		// 用户信息表、用户缓存表
	}

	// v1 短视频业务
	videoRouter := erv1.Group("/video")
	{
		// 获取推送的视频
		videoRouter.GET("/get", video.PushController)

		// 上传视频
		videoRouter.POST("/upload", video.UploadController)				// 视频上传待审核表

		// 删除视频
		videoRouter.POST("/delete", video.DeleteController)

		// 搜索视频
		videoRouter.GET("/search", video.SearchController)

		// 获取用户发布的视频
		videoRouter.GET("/users/publish", video.PublishGetController)

		// 获取用户收藏的视频
		videoRouter.GET("/users/collect", video.CollectGetController)
	}
}
