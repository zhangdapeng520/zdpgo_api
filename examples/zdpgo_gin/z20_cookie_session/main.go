package main

import (
	// 导入session包
	"github.com/gin-contrib/sessions"
	"github.com/zhangdapeng520/zdpgo_gin"

	// 导入gin框架包
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
)

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
		Session: zdpgo_gin.SessionConfig{
			OpenSession: true,
		},
	})

	g.App.GET("/test", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)

		// 通过session.Get读取session值
		// session是键值对格式数据，因此需要通过key查询数据
		if session.Get("hello") != "world" {
			// 设置session数据
			session.Set("hello", "world")
			// 删除session数据
			session.Delete("tizi365")
			// 保存session数据
			session.Save()
			// 删除整个session
			// session.Clear()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	g.Run(8080)
}
