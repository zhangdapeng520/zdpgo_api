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

	// 路由分组1
	v1 := g.App.Group("/v1")
	{
		v1.POST("/login", pong)
		v1.POST("/logout", pong)
	}

	// 路由分组2
	v2 := g.App.Group("/v2")
	{
		v2.POST("/login", pong)
		v2.POST("/logout", pong)
	}

	g.Run()
}
