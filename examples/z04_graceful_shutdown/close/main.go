package main

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 创建路由
	router := gin.Default()

	// 路径
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// 创建服务
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 退出管道
	quit := make(chan os.Signal)

	// 监听退出信号
	signal.Notify(quit, os.Interrupt)

	// 开启goroutine接收退出信号
	go func() {
		<-quit
		log.Println("接收到退出信号")
		if err := server.Close(); err != nil {
			log.Fatal("服务关闭:", err)
		}
	}()

	// 启动服务
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("服务正常退出")
		} else {
			log.Fatal("服务异常退出")
		}
	}

	log.Println("退出服务")
}
