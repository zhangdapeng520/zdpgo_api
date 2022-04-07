package main

import (
	"net/http"

	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
)

// 模拟数据库
var db = make(map[string]string)

// 设置路由
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// GET ping路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get 用户名
	r.GET("/user/:name", func(c *gin.Context) {
		// 获取路径参数
		user := c.Params.ByName("name")

		// 判断是否已存在
		value, ok := db[user]

		// 返回
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 基本的权限校验路由
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
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

	return r
}

func main() {
	r := setupRouter()
	// 监听地址 0.0.0.0:8080
	r.Run(":8080")
}
