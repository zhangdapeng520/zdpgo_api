package zdpgo_api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhangdapeng520/zdpgo_api/gin"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_password"
)

// Api API核心对象
type Api struct {
	Config *Config     // 配置对象
	App    *gin.Engine // 核心App对象
}

// New 使用默认配置创建API
func New(log *zdpgo_log.Log) *Api {
	return NewWithConfig(&Config{}, log)
}

// NewWithConfig 根据配置创建API
func NewWithConfig(config *Config, log *zdpgo_log.Log) *Api {
	a := &Api{}

	// 日志
	a.SetLog(log)

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
	if config.RateLimit == 0 {
		config.RateLimit = 3333
	}
	a.Config = config

	// App
	if config.Debug {
		a.App = gin.Default()
	} else {
		a.App = gin.Default()
		gin.SetMode(gin.ReleaseMode)
	}

	// 设置上传文件大小
	a.App.MaxMultipartMemory = config.UploadFileSize << 20

	// 加密对象
	Password = zdpgo_password.New()

	// 中间件
	if config.Middleware.Cors {
		a.AddCorsMiddleware()
	}
	if config.Middleware.RateLimit {
		a.AddRateLimitMiddleware()
	}

	// 路由
	if config.Router.HealthCheck {
		a.AddHealthCheckRouter()
	}

	// 返回对象
	return a
}

// SetLog 设置日志
func SetLog(log *zdpgo_log.Log) {
	Log = log
	gin.Log = Log
}

// SetLog 设置日志
func (a *Api) SetLog(log *zdpgo_log.Log) {
	SetLog(log)
}

// SetApp 设置APP
func (a *Api) SetApp(app *gin.Engine) {
	a.App = app
}

// TODO: 自动重启https://www.codeleading.com/article/97834445692/
// Run 运行APP
func (a *Api) Run(exitFuncList ...func()) {
	if a.App == nil {
		Log.Error("App为空，无法启动服务，请先实例化App")
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
			Log.Fatal("启动服务失败", "error", err)
		}
	}()

	// 创建退出管道
	quit := make(chan os.Signal, 1)

	// 监听退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	Log.Debug("关闭服务")

	// 优雅退出，给五秒钟响应
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Log.Fatal("服务强制关闭", "error", err)
	}

	// 执行退出函数
	for _, exitFunc := range exitFuncList {
		exitFunc()
	}

	// 完成退出
	Log.Debug("退出服务成功")
}
