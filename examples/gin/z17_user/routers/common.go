package routers

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"net/http"
	"zdpgo_gin/examples/z17_user/plugins"
	v1 "zdpgo_gin/examples/z17_user/routers/api/v1"
)

func RegisterRouter(app *gin.Engine) {
	// 挂载中间件
	// plugins.LoggerToFile() 记录日志的中间件
	app.Use(plugins.LoggerToFile())

	// 注册通用路由
	registerCommon(app)

	// 挂载版本v1的路由
	v1.RegisterV1(app)
}

func registerCommon(app *gin.Engine) {
	app.GET("/log", testLog)
	app.GET("/health", health)
}

// 健康检查
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
