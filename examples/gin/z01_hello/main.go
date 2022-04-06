package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个gin路由器
	// logger and recovery (crash-free) 中间件
	router := gin.Default()

	router.GET("/", pong)
	router.POST("/", pong)
	router.PUT("/", pong)
	router.DELETE("/", pong)
	router.PATCH("/", pong)
	router.HEAD("/", pong)
	router.OPTIONS("/", pong)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// router.Run(":3000") for a hard coded port
}
