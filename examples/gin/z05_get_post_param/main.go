package main

import (
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
)

func main() {
	router := gin.Default()

	router.POST("/test", func(c *gin.Context) {

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
	router.Run(":8080")
}
