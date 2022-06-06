package zdpgo_api

/*
@Time : 2022/6/6 15:17
@Author : 张大鹏
@File : response.go
@Software: Goland2021.3.1
@Description:
*/

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
