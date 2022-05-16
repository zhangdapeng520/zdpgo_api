// build go1.16

package main

import (
	"context"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建上下文对象，监听退出信号
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 创建路由
	router := gin.Default()

	// 监听路径
	router.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// 创建服务
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 开启协程，监听服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %s\n", err)
		}
	}()

	// 监听到退出信号
	<-ctx.Done()

	// 恢复中断信号的默认行为，并通知用户关机。
	stop()
	log.Println("优雅退出，按ctrl+c强制关闭")

	//上下文用于通知服务器，它有5秒钟的时间来完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器强制关闭: ", err)
	}

	log.Println("服务退出")
}
