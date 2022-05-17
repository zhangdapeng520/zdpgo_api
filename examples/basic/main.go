package main

import (
	"embed"
	"github.com/zhangdapeng520/zdpgo_api"
	"net/http"
)

func ping(c *zdpgo_api.Context) {
	c.JSON(http.StatusOK, zdpgo_api.JsonMap{
		"code":    10000,
		"message": "success",
	})
}

// 设置路由
func setupRouter() *zdpgo_api.Api {
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{
		Debug: true,
	})

	// 常用方法 ping路由
	api.Get("/ping", ping)
	api.Post("/ping", ping)
	api.Put("/ping", ping)
	api.Delete("/ping", ping)
	api.Patch("/ping", ping)

	// 基本的权限校验路由
	authorized := api.GetBasicAuthGroup("/", zdpgo_api.StringMap{
		"zhangdapeng": "zhangdapeng",
	})

	// POST请求
	authorized.Get("/admin", func(c *zdpgo_api.Context) {
		user := c.MustGet(zdpgo_api.AuthUserKey).(string)
		c.JSON(http.StatusOK, zdpgo_api.JsonMap{"status": user})
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
	// 监听地址 http://localhost:3333/admin
	r.Run()
}
