package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	g.App.POST("/test", func(c *gin.Context) {

		// 获取query参数
		page := c.DefaultQuery("page", "1")
		size := c.DefaultQuery("size", "20")

		// 获取form参数
		username := c.PostForm("username")
		age := c.PostForm("age")

		// 响应数据
		c.JSON(200, gin.H{
			"page":     page,
			"size":     size,
			"username": username,
			"age":      age,
		})
	})
	g.Run()
}
