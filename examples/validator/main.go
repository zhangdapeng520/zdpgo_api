package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
	"github.com/zhangdapeng520/zdpgo_api/gin"
)

type User struct {
	Name  string `form:"name" validate:"required,min=3,max=5"`
	Email string `form:"email" validate:"email"`
	Age   int8   `form:"age" validate:"gte=18,lte=20"`
}

func main() {
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{Debug: true})
	err := api.InitValidator()
	if err != nil {
		api.Log.Error("初始化校验器失败", "error", err)
		return
	}

	api.App.GET("/language", func(context *gin.Context) {
		var user User
		// 绑定查询参数
		err = context.ShouldBindQuery(&user)
		if err != nil {
			context.JSON(500, gin.H{"msg": err})
			return
		}

		// 使用验证器验证
		errData := api.Validate(user)
		if errData != nil {
			context.JSON(500, gin.H{
				"errData": errData,
			})
			return
		}
		context.JSON(200, gin.H{"msg": "校验成功"})
	})

	// 测试失败：http://127.0.0.1:3333/language
	// 测试成功：http://127.0.0.1:3333/language?name=abcd&age=19&email=xxx@example.com
	api.Run()
}
