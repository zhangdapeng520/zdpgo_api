# zdpgo_api
基于gin二次封装的一个后端api快速开发框架

项目地址：https://github.com/zhangdapeng520/zdpgo_api

## 版本历史
- v1.0.0  2022/2/9
- v1.0.1  2022/2/11  指定默认模板目录template和默认静态文件夹目录static
- v1.0.2  2022/2/11  移除本地依赖
- v1.0.3  2022/2/11  新增日志中间件和recover中间件
- v1.0.4  2022/3/8   优化模板和静态目录
- v1.0.5  2022/3/8   增加常用的统一返回对象
- v1.0.6  2022/3/8   将挂载通用路由设置为开关量
- v1.0.7  2022/3/10  支持配置服务相关参数
- v1.0.8  2022/3/10  支持viper读取和设置配置
- v1.0.9  2022/3/11  logger日志支持配置
- v1.1.0  2022/3/12  修复日志不打印body
- v1.1.1  2022/4/8   新增详细日志
- v1.1.2  2022/4/11  详细日志升级，支持查看的具体Body中的JSON数据
- v1.1.3  2022/4/21  文件上传
- v1.1.4  2022/5/16  新增：Api核心对象
- v1.1.5  2022/5/17  优化：日志中间件优化
- v1.1.6  2022/5/17  新增：添加静态文件服务

## 示例
### 基本用法
```go
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
```

### 基本权限校验
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
	"github.com/zhangdapeng520/zdpgo_api/gin"
)

// 用户列表
var secrets = gin.H{
	"zhangdapeng": gin.H{"email": "zhangdapeng@qq.com", "phone": "123433"},
}

func getSecrets(c *gin.Context) {
	// 获取用户名
	user := c.MustGet(gin.AuthUserKey).(string)

	// 取出数据
	if secret, ok := secrets[user]; ok {
		c.JSON(200, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(200, gin.H{"user": user, "secret": "No SECRET :("})
	}
}

func main() {
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{Debug: true})

	// 创建权限校验路由分组
	authorized := api.GetBasicAuthGroup("/admin", map[string]string{"zhangdapeng": "zhangdapeng"})

	// 校验接口
	authorized.GET("/secrets", getSecrets)

	// 启动服务。访问：http://localhost:3333/admin/secrets
	api.Run()
}
```

### 文件上传
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
)

func main() {
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{Debug: true})
	api.AddUploadRouter("/upload", "file", "uploads")
	api.Run()
}
```