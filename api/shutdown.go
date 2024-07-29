package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 天然支持优雅退出的一种启动方式
// @param addr 服务地址
// @param router 路由对象，调用 api.NewRouter() 生成
func Run(addr string, router http.Handler) {
	RunWith(addr, router)
}

// RunWith 天然支持优雅退出的一种启动方式
// @param addr 服务地址
// @param router 路由对象，调用 api.NewRouter() 生成
// @param funcs 被注册的退出方法，当监听到退出信号的时候自动执行
func RunWith(addr string, router http.Handler, funcs ...func()) {
	// 创建服务
	server := &http.Server{Addr: addr, Handler: router}
	log.Printf("采用优雅退出的方式启动服务，该HTTP服务将在 %s 启动，当监听到退出信号时会主动释放资源后再退出\n", addr)

	// 创建服务的上下文
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// 监听退出的信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		log.Println("监听到退出信号，等待服务器关闭，30秒后将强制关闭服务")

		// 执行退出方法
		log.Println("执行被注册的退出方法")
		for _, f := range funcs {
			f()
		}

		// 关闭信号，宽限期为 30 秒
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("优雅关机超时......强制退出。")
			}
		}()

		// 触发优雅关机
		log.Println("等待其他服务器资源关闭...")
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// 运行服务器
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// 等待服务器上下文停止
	<-serverCtx.Done()
}
