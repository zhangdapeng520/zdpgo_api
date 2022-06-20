package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
	"github.com/zhangdapeng520/zdpgo_log"
)

/*
@Time : 2022/6/6 15:14
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	api := zdpgo_api.New(zdpgo_log.Tmp)

	api.Post("/aes", func(ctx *zdpgo_api.Context) {
		// 解析json数据
		var jsonData struct {
			Username string `json:"username"`
			Age      int    `json:"age"`
		}
		err := ctx.GetAesTextBodyToJson(&jsonData)
		if err != nil {
			panic(err)
		}

		// 加密响应数据
		response := &zdpgo_api.Response{
			Code:   10000,
			Msg:    "success",
			Status: true,
			Data:   jsonData,
		}
		ctx.ResponseAesStr(response)
	})

	// 启动
	api.Run()
}
