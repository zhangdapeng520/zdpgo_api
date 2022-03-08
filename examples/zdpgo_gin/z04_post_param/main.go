package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 获取POST参数
	g.App.POST("/test", func(c *gin.Context) {
		username := c.PostForm("username")
		age := c.DefaultPostForm("age", "22") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":   true,
			"username": username,
			"age":      age,
		})
	})
	g.Run()
}
