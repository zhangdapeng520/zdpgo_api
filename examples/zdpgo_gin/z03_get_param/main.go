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

	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	g.App.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("first_name", "张大鹏")
		lastname := c.Query("last_name") // 是 c.Request.URL.Query().Get("lastname") 的简写

		c.String(http.StatusOK, "你好 %s %s", firstname, lastname)
	})
	//g.Run(8080)
	g.Run()
}
