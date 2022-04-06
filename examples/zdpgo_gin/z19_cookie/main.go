package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
)

func readCookie(c *gin.Context) {
	// 根据cookie名字读取cookie值
	data, err := c.Cookie("token")
	if err != nil {
		// 直接返回cookie值
		c.JSON(200, gin.H{
			"status": false,
			"msg":    "error",
			"code":   10000,
		})
		return
	}
	c.JSON(200, gin.H{
		"cookie": data,
		"status": true,
		"msg":    "success",
		"code":   10000,
	})
}

func deleteCookie(c *gin.Context) {
	// 设置cookie  MaxAge设置为-1，表示删除cookie
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"status": true,
		"msg":    "success",
		"code":   10000,
	})
}

func setCookie(c *gin.Context) {
	// 设置cookie
	c.SetCookie("token", "abc_token_zhangdapeng", 3600, "/", "*", false, false)
	c.JSON(200, gin.H{
		"status": true,
		"msg":    "success",
		"code":   10000,
	})
}

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	g.App.GET("/cookie", readCookie)      // 获取cookie
	g.App.DELETE("/cookie", deleteCookie) // 删除cookie
	g.App.POST("/cookie", setCookie)      // 新增cookie
	
	g.Run(8888)
}
