package main

import (
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// 路由1：注意，*gin.Engine可以被当成http.Handler使用
func router01() http.Handler {
	// 创建服务
	e := gin.New()
	e.Use(gin.Recovery())

	// 监听路径
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

// 路由2
func router02() http.Handler {
	// 创建服务
	e := gin.New()
	e.Use(gin.Recovery())

	// 监听路径
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	// 创建服务1: http://localhost:8080
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 创建服务2: http://localhost:8081
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 使用错误组启动goroutine
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	// 如果都出错了，则退出
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
