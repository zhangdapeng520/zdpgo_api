package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// simulate some private data
var secrets = gin.H{
	"zhangdapeng": gin.H{"email": "zhangdapeng@bar.com", "phone": "123433"},
	"austin":      gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":        gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// 使用BasicAuth中间件，传入用户字典
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"zhangdapeng": "zhangdapeng", // 用户名和密码
		"austin":      "1234",
		"lena":        "hello2",
		"manu":        "4321",
	}))

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

	r.Run(":8080")
}
