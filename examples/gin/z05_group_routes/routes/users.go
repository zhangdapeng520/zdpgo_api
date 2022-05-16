package routes

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"net/http"
)

// 添加user分组路由
func addUserRoutes(rg *gin.RouterGroup) {

	// 创建users分组
	users := rg.Group("/users")

	// 添加users分组下的路由
	users.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "users")
		return
	})
	users.GET("/comments/", func(c *gin.Context) {
		c.String(http.StatusOK, "users comments")
	})
	users.GET("/pictures/", func(c *gin.Context) {
		c.String(http.StatusOK, "users pictures")
	})
}
