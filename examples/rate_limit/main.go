package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
)

func main() {
	api := zdpgo_api.NewWithConfig(&zdpgo_api.Config{
		Debug:     true,
		Port:      3333,
		RateLimit: 3333,
		Middleware: zdpgo_api.MiddlewareConfig{
			RateLimit: true,
		},
		Router: zdpgo_api.RouterConfig{
			HealthCheck: true, // 健康检查
		},
	})
	api.AddRateLimitMiddleware()
	api.Post("/aes", func(ctx *zdpgo_api.Context) {
		var jsonData struct {
			Username string `json:"username"`
			Age      int    `json:"age"`
		}

		// 加密响应数据
		response := &zdpgo_api.Response{
			Code:   10000,
			Msg:    "success",
			Status: true,
			Data:   jsonData,
		}
		ctx.JSON(200, ctx.GetResponseSuccess(response))
	})

	// 启动
	api.Run()
}
