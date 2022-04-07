package main

import (
	"context"
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建路由
	router := gin.Default()

	// 监听路径
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// 创建服务
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 开启协程，启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务失败: %s\n", err)
		}
	}()

	// 创建退出管道
	quit := make(chan os.Signal, 1)

	// 监听退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭服务。。。")

	// 优雅退出，给五秒钟响应
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务强制关闭: ", err)
	}

	log.Println("服务退出")
}
