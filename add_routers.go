package zdpgo_api

/*
@Time : 2022/6/27 11:02
@Author : 张大鹏
@File : add_routers.go
@Software: Goland2021.3.1
@Description:
*/

// AddHealthCheckRouter 添加健康检查接口
func (a *Api) AddHealthCheckRouter() {
	a.Get("/health", func(ctx *Context) {
		success := ctx.GetResponseSuccess(nil)
		ctx.JSON(200, success)
	})
}
