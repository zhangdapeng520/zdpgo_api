package routes

import (
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
	"net/http"
)

// 添加路由
func addPingRoutes(rg *gin.RouterGroup) {
	// 创建ping分组
	ping := rg.Group("/ping")

	// 添加ping分组下的路径
	ping.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
