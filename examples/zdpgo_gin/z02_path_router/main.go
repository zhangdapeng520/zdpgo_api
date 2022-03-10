package main

import (
	"github.com/zhangdapeng520/zdpgo_gin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 创建app
	app := g.CreateApp()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	app.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "你的名字是 %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	app.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " 要 " + action
		c.String(http.StatusOK, message)
	})

	g.Run()
}
