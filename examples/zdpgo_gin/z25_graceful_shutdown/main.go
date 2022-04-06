package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_gin"
	"net/http"
	"time"

	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
)

func shutDown1() {
	fmt.Println("退出程序时执行的方法1")
}

func shutDown2() {
	fmt.Println("退出程序时执行的方法2")
}

func main() {
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})
	g.App.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	g.Run(shutDown1, shutDown2)
}
