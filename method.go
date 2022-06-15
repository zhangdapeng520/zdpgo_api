package zdpgo_api

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"strings"
	"unsafe"
)

/*
@Time : 2022/5/17 16:01
@Author : 张大鹏
@File : method.go
@Software: Goland2021.3.1
@Description: method方法相关
*/

// Any 处理任意方法类型
func (a *Api) Any(method, routerPath string, handleFuncList ...func(ctx *Context)) {

	// 异常情况
	if handleFuncList == nil || len(handleFuncList) == 0 {
		Log.Error("API处理方法不能为空", "handleFuncList", handleFuncList)
		return
	}

	// 处理结果的方法对象
	handleFuncObj := func(ctxGin *gin.Context) {
		for _, handleFunc := range handleFuncList {
			ctxApiPointer := unsafe.Pointer(ctxGin) // 将父类实例转为通用指针
			ctxApi := (*Context)(ctxApiPointer)     //再转换为子类结构体
			handleFunc(ctxApi)
		}
	}

	switch strings.ToUpper(method) {
	case "GET":
		a.App.GET(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "POST":
		a.App.POST(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "PUT":
		a.App.PUT(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "DELETE":
		a.App.DELETE(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "PATCH":
		a.App.PATCH(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "HEAD":
		a.App.HEAD(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	default:
		Log.Error("暂不支持此类型的方法", "method", method)
	}
}

// Get 处理GET方法类型
func (a *Api) Get(routerPath string, handleFuncList ...func(ctx *Context)) {
	a.Any("GET", routerPath, handleFuncList...)
}

// Post 处理POST方法类型
func (a *Api) Post(routerPath string, handleFuncList ...func(ctx *Context)) {
	a.Any("POST", routerPath, handleFuncList...)
}

// Delete 处理DELETE方法类型
func (a *Api) Delete(routerPath string, handleFuncList ...func(ctx *Context)) {
	a.Any("DELETE", routerPath, handleFuncList...)
}

// Put 处理PUT方法类型
func (a *Api) Put(routerPath string, handleFuncList ...func(ctx *Context)) {
	a.Any("PUT", routerPath, handleFuncList...)
}

// Patch 处理PATCH方法类型
func (a *Api) Patch(routerPath string, handleFuncList ...func(ctx *Context)) {
	a.Any("PATCH", routerPath, handleFuncList...)
}
