package zdpgo_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterCommonRouter 注册通用路由
func (g *Gin) RegisterCommonRouter(app *gin.Engine) {
	app.GET("/health", g.GetRouterCommonHealth)   // 监控检查
	app.GET("/captcha", g.GetRouterCommonCaptcha) // 图片验证码
}

// GetRouterCommonHealth 健康检查地址
func (g *Gin) GetRouterCommonHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
	})
}
