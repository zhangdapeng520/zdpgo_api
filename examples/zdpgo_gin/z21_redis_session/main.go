package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
		Session: zdpgo_gin.SessionConfig{
			OpenSession: true,
			SessionType: "redis",
		},
	})

	g.App.GET("/test", func(c *gin.Context) {
		// 创建session
		session := sessions.Default(c)

		// 设置session
		session.Set("count", 33)

		// 获取session
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)

		// 保存session
		err := session.Save()
		if err != nil {
			fmt.Println("保存session失败：", err.Error())
		}
		c.JSON(200, gin.H{"count": count})
	})
	g.Run(8080)
}
