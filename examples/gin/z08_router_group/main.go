package main

import (
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	router := gin.Default()

	// 路由分组1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", pong)
		v1.POST("/logout", pong)
	}

	// 路由分组2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", pong)
		v2.POST("/logout", pong)
	}

	router.Run(":8080")
}
