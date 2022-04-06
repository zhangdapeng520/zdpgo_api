package zdpgo_gin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
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
func (g *Gin) Run(shutdownFuncArr ...func()) {
	// 默认端口
	if g.config.Server.Host == "" {
		g.config.Server.Host = "0.0.0.0"
	}
	if g.config.Server.Port == 0 {
		g.config.Server.Port = 8888
	}
	if g.config.Server.ReadTimeout == 0 {
		g.config.Server.ReadTimeout = 33
	}
	if g.config.Server.WriteTimeout == 0 {
		g.config.Server.WriteTimeout = 33
	}
	if g.config.Server.MaxHeaderBytes == 0 {
		g.config.Server.MaxHeaderBytes = 1024 * 1024
	}

	// 创建应用
	address := fmt.Sprintf(
		"%s:%d",
		g.config.Server.Host,
		g.config.Server.Port,
	)

	srv := &http.Server{
		Addr:           address,
		Handler:        g.App,
		ReadTimeout:    time.Duration(g.config.Server.ReadTimeout) * time.Second,  // 读超时时间
		WriteTimeout:   time.Duration(g.config.Server.WriteTimeout) * time.Second, // 写超时时间
		MaxHeaderBytes: int(g.config.Server.MaxHeaderBytes),                       // 请求头大小限制
	}

	// 启动服务
	g.log.Info("启动服务器", "host", g.config.Server.Host, "port", g.config.Server.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			g.log.Error("服务退出", "error", err.Error())
		}
	}()

	// 创建退出信号管道
	quit := make(chan os.Signal)

	// 监听信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	g.log.Info("关闭服务...")

	// 设置定时器，关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 执行退出函数
	for _, shutdownFunc := range shutdownFuncArr {
		shutdownFunc()
	}

	// 关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		g.log.Fatal("强制关闭服务", "error", err.Error())
	}

	// 正常关闭服务
	g.log.Info("关闭服务成功！")
}
