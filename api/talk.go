package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有跨域请求，生产环境中应根据实际情况配置
	},
}

func InitTalkRouter(r *gin.RouterGroup) {
	// WebSocket 连接路由
	r.GET("/ws", WebSocketConnect)
}

// WebSocketConnect 连接处理函数
func WebSocketConnect(c *gin.Context) {
	// 升级 HTTP 请求为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade Error:", err)
		return
	}
	defer conn.Close()

	// 处理 WebSocket 连接
	for {
		// 读取消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("ReadMessage Error:", err)
			return
		}

		// 打印收到的消息
		fmt.Printf("Received: %s", p)

		// 发送消息回客户端
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println("WriteMessage Error:", err)
			return
		}
	}
}
