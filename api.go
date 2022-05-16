package zdpgo_api

import (
	"context"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "github.com/zhangdapeng520/zdpgo_log"

// Api API核心对象
type Api struct {
	Config *Config        // 配置对象
	Log    *zdpgo_log.Log // 日志对象
	App    *gin.Engine    // 核心App对象
}

// New 使用默认配置创建API
func New() *Api {
	return NewWithConfig(Config{})
}

// NewWithConfig 根据配置创建API
func NewWithConfig(config Config) *Api {
	a := &Api{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_api.log"
	}
	a.Log = zdpgo_log.NewWithDebug(config.Debug, config.LogFilePath)

	// 配置
	if config.Host == "" {
		config.Host = "0.0.0.0"
	}
	if config.Port == 0 {
		config.Port = 3333
	}
	if config.UploadFileSize == 0 {
		config.UploadFileSize = 33
	}
	a.Config = &config

	// App
	if config.Debug {
		a.App = NewGinWithLog(true)
	} else {
		a.App = gin.Default()
		gin.SetMode(gin.ReleaseMode)
	}

	// 设置上传文件大小
	a.App.MaxMultipartMemory = config.UploadFileSize << 20

	// 返回对象
	return a
}

// SetApp 设置APP
func (a *Api) SetApp(app *gin.Engine) {
	a.App = app
}

// Run 运行APP
func (a *Api) Run(exitFuncList ...func()) {
	if a.App == nil {
		a.Log.Error("App为空，无法启动服务，请先实例化App")
		return
	}

	// 创建服务
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.Config.Host, a.Config.Port),
		Handler: a.App,
	}

	// 开启协程，启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Log.Fatal("启动服务失败", "error", err)
		}
	}()

	// 创建退出管道
	quit := make(chan os.Signal, 1)

	// 监听退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Log.Debug("关闭服务")

	// 优雅退出，给五秒钟响应
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		a.Log.Fatal("服务强制关闭", "error", err)
	}

	// 执行退出函数
	for _, exitFunc := range exitFuncList {
		exitFunc()
	}

	// 完成退出
	a.Log.Debug("退出服务成功")
}

// NewGinWithLog 根据是否使用详细日志创建Gin的实例
// @param isDetailLogger 是否使用详细日志，开启后，会记录请求头，form，body的详细信息
func NewGinWithLog(isDetailLogger bool) *gin.Engine {
	// 创建gin对象
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		IsDetailLogger: isDetailLogger, // 使用详细日志
	}), gin.Recovery())

	// 设置模式
	if isDetailLogger {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 返回实例
	return r
}
