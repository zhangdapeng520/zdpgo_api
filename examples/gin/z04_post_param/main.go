package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
)

func main() {
	router := gin.Default()

	// 获取POST参数
	router.POST("/test", func(c *gin.Context) {
		username := c.PostForm("username")
		age := c.DefaultPostForm("age", "22") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":   true,
			"username": username,
			"age":      age,
		})
	})
	router.Run(":8080")
}
