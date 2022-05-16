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
