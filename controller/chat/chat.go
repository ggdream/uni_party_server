package chat

import (
	"fmt"
	"gateway/middleware"
	"gateway/model/chat"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	websocketx "gateway/tools/websocket"
)

var (
	// 协议转换器
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 封装的WS-Handler
	socketer = websocketx.New()
)

// ChatHandler 聊天入口
func ChatHandler(c *gin.Context) {
	// 获取GET请求的query参数
	var queryModel chat.Transform
	if err := c.ShouldBindQuery(&queryModel); err != nil {
		// TODO: 返回错误自定义状态码
		return
	}

	uid := c.GetUint(middleware.KeyUID)

	// 尝试协议转换，并获取WS套接字`conn`
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			// TODO: 日志收集
			fmt.Println(err)
		}
	}(conn)

	// 登陆和退出
	socketer.Join(uid, conn)
	defer socketer.Exit(uid)

	// 存储该套接字
	// 循环读取本台机器上套接字的接收
	// 打开Redis连接，发布订阅，监听消息
	// 不使用消息队列是因为消息要发送至每台，让服务器自己做判断
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			// 读取有问题，就断开与客户端的连接
			return
		}

		wrapper, err := websocketx.NewWrapper(data)
		if err != nil {
			// 数据格式不符合预设规则
			continue
		}

		if err := socketer.Handle(wrapper); err != nil {
			fmt.Println(err)
		}
	}
}
