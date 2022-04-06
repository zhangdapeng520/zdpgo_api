package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
	"net/http"
)

// 私密的数据
var secrets = gin.H{
	"zhangdapeng": gin.H{"email": "zhangdapeng@bar.com", "phone": "123433"},
}

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug:            true,
		OpenCommonRouter: true, // 开启通用路由
	})

	// 对指定分组使用BasicAuth中间件，传入用户字典
	authorized := g.App.Group("/admin")
	g.OpenBasicAuth(map[string]string{ // 开启BasicAuth校验，指定用户名和密码
		"zhangdapeng": "zhangdapeng",
	}, authorized)

	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户名，被BasicAuth中间件提供
		user := c.MustGet(gin.AuthUserKey).(string)

		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	g.Run()
}
