package zdpgo_api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Api API核心对象
type Api struct {
	Config *Config // 配置对象
	App    *Engine // 核心App对象
}

// New 使用默认配置创建API
func NewApi() *Api {
	return NewWithConfig(&Config{})
}

// NewWithConfig 根据配置创建API
func NewWithConfig(config *Config) *Api {
	a := &Api{}

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
		a.App = Default()
	} else {
		a.App = Default()
		SetMode(ReleaseMode)
	}

	// 设置上传文件大小
	a.App.MaxMultipartMemory = config.UploadFileSize << 20

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
	if config.Router.Static {
		a.AddStaticRouter("/static", "./static")
	}
	if config.Router.Upload {
		a.AddUploadRouter("/upload", "file", "./uploads")
	}

	// 返回对象
	return a
}

// SetApp 设置APP
func (a *Api) SetApp(app *Engine) {
	a.App = app
}

// TODO: 自动重启https://www.codeleading.com/article/97834445692/
// Run 运行APP
func (a *Api) Run(exitFuncList ...func()) {
	if a.App == nil {
		return
	}

	// 创建服务
	addr := fmt.Sprintf("%s:%d", a.Config.Host, a.Config.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: a.App,
	}

	// 开启协程，启动服务
	fmt.Println("启动REST API服务：", addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// 创建退出管道
	quit := make(chan os.Signal, 1)

	// 监听退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅退出，给五秒钟响应
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	// 执行退出函数
	for _, exitFunc := range exitFuncList {
		exitFunc()
	}

	// 完成退出
	fmt.Println("退出服务成功")
}
