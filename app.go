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
func (g *Gin) Run(port ...uint16) {
	// 默认端口
	var p uint16 = 8080

	// 使用传进来的端口
	if len(port) > 0 {
		p = port[0]
	}
	g.log.Info("启动服务器", "端口", p)
	if err := g.App.Run(fmt.Sprintf(":%d", p)); err != nil {
		g.log.Panic("启动失败", "error", err.Error())
	}
}
