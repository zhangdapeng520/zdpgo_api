package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
	"log"
	"time"
)

func main() {
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
		Server: zdpgo_gin.ServerConfig{
			Port: 8080,
		},
	})

	g.App.GET("/long_async", func(c *gin.Context) {
		// 复制内部的goroutine
		cCp := c.Copy()

		// 异步执行方法
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path) // 使用ctx上下文
		}()
		c.JSON(200, gin.H{
			"path": c.Request.URL.Path,
		})
	})

	g.App.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + c.Request.URL.Path)
		c.JSON(200, gin.H{
			"path": c.Request.URL.Path,
		})
	})

	g.Run()
}
