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

	// 全局中间件
	// 使用 Logger 中间件
	g.App.Use(gin.Logger())

	// 使用 Recovery 中间件
	g.App.Use(gin.Recovery())

	// 路由添加中间件，可以添加任意多个
	g.App.GET("/test", gin.Logger(), gin.Recovery(), pong)

	authorized := g.App.Group("/")

	// 路由组中添加中间件
	authorized.Use(gin.Recovery())
	{
		authorized.POST("/login", pong)
		authorized.POST("/logout", pong)

		// 嵌套分组
		testing := authorized.Group("test1")
		testing.POST("/abc", pong)
	}

	// 启动服务
	g.Run()
}
