package main

import (
	"embed"
	"github.com/zhangdapeng520/zdpgo_api"
	"net/http"

	"github.com/zhangdapeng520/zdpgo_api/gin"
)

// 模拟数据库
var db = make(map[string]string)

// 设置路由
func setupRouter() *zdpgo_api.Api {
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{
		Debug: true,
	})

	// GET ping路由
	api.App.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get 用户名
	api.App.GET("/user/:name", func(c *gin.Context) {
		// 获取路径参数
		user := c.Params.ByName("name")

		// 返回
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	// 基本的权限校验路由
	authorized := api.App.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	// POST请求
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return api
}

//go:embed uploads downloads
var fsObj embed.FS

func main() {
	r := setupRouter()

	// 添加静态目录
	r.AddStaticRouter("/static", "./uploads")
	r.AddStaticFSRouter("/fs", fsObj)

	// 监听地址 http://localhost:3333/ping?a=111&b=222#abc
	// 监听地址 http://localhost:3333/user/zhangdapeng
	// 监听地址 http://localhost:3333/static/test1.jpg
	// 监听地址 http://localhost:3333/fs/uploads/test1.jpg
	// 监听地址 http://localhost:3333/fs/downloads/test1.jpg
	r.Run()
}
