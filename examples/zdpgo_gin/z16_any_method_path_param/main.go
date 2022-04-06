package main

import (
	"github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"log"
	"net/http"
)

func pong(c *gin.Context) {
	json := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "success",
		"data": json,
	})
}

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 定义任意方法，任意路径，任意body的路由

	g.App.GET("/*path", pong)
	g.App.POST("/*path", pong)
	g.App.PUT("/*path", pong)
	g.App.DELETE("/*path", pong)
	g.App.PATCH("/*path", pong)

	g.Run(8888)
}
