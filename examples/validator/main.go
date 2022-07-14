package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
)

type User struct {
	Name  string `form:"name" validate:"required,min=3,max=5"`
	Email string `form:"email" validate:"email"`
	Age   int8   `form:"age" validate:"gte=18,lte=20"`
}

func main() {
	api := zdpgo_api.NewApi()
	err := api.InitValidator()
	if err != nil {
		return
	}

	api.App.GET("/language", func(context *zdpgo_api.Context) {
		var user User
		// 绑定查询参数
		err = context.ShouldBindQuery(&user)
		if err != nil {
			context.JSON(500, zdpgo_api.H{"msg": err})
			return
		}

		// 使用验证器验证
		errData := api.Validate(user)
		if errData != nil {
			context.JSON(500, zdpgo_api.H{
				"errData": errData,
			})
			return
		}
		context.JSON(200, zdpgo_api.H{"msg": "校验成功"})
	})

	// 测试失败：http://127.0.0.1:3333/language
	// 测试成功：http://127.0.0.1:3333/language?name=abcd&age=19&email=xxx@example.com
	api.Run()
}
