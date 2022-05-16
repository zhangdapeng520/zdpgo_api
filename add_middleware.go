package zdpgo_api

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"net/http"
)

/*
@Time : 2022/5/16 17:51
@Author : 张大鹏
@File : add_middleware.go
@Software: Goland2021.3.1
@Description: 添加中间件
*/

// AddCorsMiddleware 添加跨域中间件
func (a *Api) AddCorsMiddleware() {
	// 跨域中间件函数
	coreMiddleware := func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token, zdp-token, ZDPToken")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}

	// 添加快鱼中间件
	a.App.Use(coreMiddleware)
}
