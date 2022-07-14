package zdpgo_api

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

// 获取成功的响应
func (c *Context) GetResponseSuccess(data interface{}) Response {
	return Response{
		Code:   10000,
		Msg:    "success",
		Status: true,
		Data:   data,
	}
}
