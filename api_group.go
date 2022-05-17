package zdpgo_api

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"strings"
	"unsafe"
)

/*
@Time : 2022/5/17 17:52
@Author : 张大鹏
@File : api_group.go
@Software: Goland2021.3.1
@Description: api_group API分组相关
*/

// ApiGroup API分组
type ApiGroup struct {
	gin.RouterGroup
}

func (g *ApiGroup) Any(method, routerPath string, handleFuncList ...func(ctx *Context)) {

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
		g.GET(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "POST":
		g.POST(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "PUT":
		g.PUT(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "DELETE":
		g.DELETE(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "PATCH":
		g.PATCH(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	case "HEAD":
		g.HEAD(routerPath, func(ctxGin *gin.Context) {
			handleFuncObj(ctxGin)
		})
	default:
		Log.Error("暂不支持此类型的方法", "method", method)
	}
}

// Get 处理GET方法类型
func (g *ApiGroup) Get(routerPath string, handleFuncList ...func(ctx *Context)) {
	g.Any("GET", routerPath, handleFuncList...)
}

// Post 处理POST方法类型
func (g *ApiGroup) Post(routerPath string, handleFuncList ...func(ctx *Context)) {
	g.Any("POST", routerPath, handleFuncList...)
}

// Delete 处理DELETE方法类型
func (g *ApiGroup) Delete(routerPath string, handleFuncList ...func(ctx *Context)) {
	g.Any("DELETE", routerPath, handleFuncList...)
}

// Put 处理PUT方法类型
func (g *ApiGroup) Put(routerPath string, handleFuncList ...func(ctx *Context)) {
	g.Any("PUT", routerPath, handleFuncList...)
}

// Patch 处理PATCH方法类型
func (g *ApiGroup) Patch(routerPath string, handleFuncList ...func(ctx *Context)) {
	g.Any("PATCH", routerPath, handleFuncList...)
}
