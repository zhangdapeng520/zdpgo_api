package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

var (
	g *zdpgo_gin.Gin
)

func init() {
	g = zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug:         true,
		OpenWebsocket: true, // 开启websocket
	})

}

//webSocket请求ping 返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := g.CreateWebsocket(c)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	g.App.GET("/ping", ping)
	g.Run(2303)
}
