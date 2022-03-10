package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 创建app
	app := g.CreateApp()

	// 挂载路由
	app.GET("/", pong)
	app.POST("/", pong)
	app.PUT("/", pong)
	app.DELETE("/", pong)
	app.PATCH("/", pong)
	app.HEAD("/", pong)
	app.OPTIONS("/", pong)

	// 启动服务
	g.Run()
}
