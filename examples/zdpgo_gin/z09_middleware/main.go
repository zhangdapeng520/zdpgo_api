package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func SetupRouter() *zdpgo_gin.Gin {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
		Server: zdpgo_gin.ServerConfig{
			Records: []string{"header", "body", "url", "form"},
		},
	})
	return g
}

func main() {
	g := SetupRouter()

	authorized := g.App.Group("/")

	// 路由组中添加中间件
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
