package websocket

import (
	"errors"
	"gateway/mongo"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"gateway/tools/cmap"
)

type (
	// cmap的内层interface{}
	// deviceConnMapType = map[DeviceType]*websocket.Conn

	// WS处理器
	WebSocket struct {
		// 存储用户在各个平台的WebSocket套接字
		// key: uid，value: map[Enum]*websocket.Conn
		conn *cmap.CMap
	}
)

func New() *WebSocket {
	return &WebSocket{
		conn: cmap.New(64),
	}
}

// Join 用户登录
func (w *WebSocket) Join(uid uint, conn *websocket.Conn) {
	w.conn.Set(strconv.Itoa(int(uid)), conn)
}

// Exit 用户退出
func (w *WebSocket) Exit(uid uint) {
	w.conn.Del(strconv.Itoa(int(uid)))
}

// Handle 对外处理函数
func (w *WebSocket) Handle(wrapper *Wrapper) error {
	return w.handle(wrapper)
}

// handle 处理接收到的消息
func (w *WebSocket) handle(wrapper *Wrapper) error {
	// switch NewChatType(wrapper.Type) {
	// case TEXT:
	// case IMAGE, AUDIO, VIDEO:
	// default:
	// }

	// 封装响应消息体
	datetime := time.Now().UnixNano()
	wrapper.Modify(datetime, "xxxx")

	document := mongo.ChatDocument{
		Type:     wrapper.Type,
		FromUID:  wrapper.FromUID,
		ToUID:    wrapper.ToUID,
		GroupID:  wrapper.GroupID,
		Message:  wrapper.Message,
		URL:      wrapper.URL,
		Others:   wrapper.Others,
		MID:      wrapper.MID,
		Datetime: wrapper.Datetime,
		Version:  wrapper.Version,
	}
	if err := document.Insert(); err != nil {
		return errors.New("insert mongo failed")
	}

	toUID := strconv.Itoa(wrapper.ToUID)
	// 获取WebSocket套接字
	conn, ok := w.conn.Get(toUID)
	if !ok {
		// 用户未登录则直接返回
		return errors.New("user no login")
	}

	// 把数据发送至目的地
	if err := conn.(*websocket.Conn).WriteJSON(wrapper); err != nil {
		w.conn.Del(toUID)
		return conn.(*websocket.Conn).Close()
	}

	return nil
}
