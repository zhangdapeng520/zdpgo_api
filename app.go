package zdpgo_gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateApp 创建app
func (g *Gin) CreateApp() *gin.Engine {
	app := gin.Default()

	// 健康检查地址
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域
	app.Use(g.MiddlewareCors())

	// 挂载通用路由
	app.GET("/captcha", g.GetRouterCommonCaptcha)

	// 初始化
	g.App = app

	// 返回
	return app
}

// Run 启动服务
func (g *Gin) Run(port uint16) {
	g.log.Info("启动服务器", "端口", port)
	if err := g.App.Run(fmt.Sprintf(":%d", port)); err != nil {
		g.log.Panic("启动失败", "error", err.Error())
	}
}
