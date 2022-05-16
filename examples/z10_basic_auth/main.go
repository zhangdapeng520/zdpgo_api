package main

import (
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
	r := gin.Default()

	// 创建路由分组，并使用基础权限校验中间件
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"zhangdapeng": "zhangdapeng", //用户名：密码
	}))

	// 校验接口
	authorized.GET("/secrets", getSecrets)

	// 启动服务。访问：http://localhost:8080/admin/secrets
	r.Run(":8080")
}
