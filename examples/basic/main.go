package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zhangdapeng520/zdpgo_api"
)

func ping(c *zdpgo_api.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	c.JSON(http.StatusOK, zdpgo_api.JsonMap{
		"code":    10000,
		"message": "success111",
		"data":    string(data),
	})
}

func form(c *zdpgo_api.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	response := c.GetResponseSuccess(map[string]interface{}{
		"username": username,
		"password": password,
	})
	c.JSON(http.StatusOK, response)
}

func longTime(c *zdpgo_api.Context) {
	time.Sleep(time.Second * 30) // 30秒
	c.JSON(http.StatusOK, zdpgo_api.JsonMap{
		"code":    10000,
		"message": "success",
	})
}

func jsonRouter(c *zdpgo_api.Context) {
	jsonData := make(map[string]interface{}) //注意该结构接受的内容
	_ = c.BindJSON(&jsonData)
	response := c.GetResponseSuccess(jsonData)
	c.JSON(http.StatusOK, response)
}

func textRouter(c *zdpgo_api.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("body", string(body))
	c.String(http.StatusOK, string(body))
}

// 设置路由
func setupRouter() *zdpgo_api.Api {
	api := zdpgo_api.NewApi()

	// 常用方法 ping路由
	api.Get("/ping", ping)
	api.Post("/ping", ping)
	api.Put("/ping", ping)
	api.Delete("/ping", ping)
	api.Patch("/ping", ping)

	// 表单方法
	api.Get("/form", form)
	api.Post("/form", form)
	api.Put("/form", form)
	api.Delete("/form", form)
	api.Patch("/form", form)

	// 耗时很长的方法
	api.Get("/long", longTime)

	// 获取json数据并返回的方法
	api.Post("/json", jsonRouter)
	api.Put("/json", jsonRouter)
	api.Delete("/json", jsonRouter)
	api.Patch("/json", jsonRouter)

	// 获取文本数据并返回
	api.Post("/text", textRouter)

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

//go:embed downloads
var fsObj embed.FS

func main() {
	api := setupRouter()

	// 添加静态目录
	api.AddStaticRouter("/static", "./uploads")
	api.AddStaticFSRouter("/fs", fsObj)

	// 添加文件上传
	api.AddUploadRouter("/upload", "file", "uploads")

	// 监听地址 http://localhost:3333/ping?a=111&b=222#abc
	// 监听地址 http://localhost:3333/user/zhangdapeng
	// 监听地址 http://localhost:3333/static/test1.jpg
	// 监听地址 http://localhost:3333/fs/uploads/test1.jpg
	// 监听地址 http://localhost:3333/fs/downloads/test1.jpg
	// 监听地址 http://localhost:3333/admin
	api.Run()
}
