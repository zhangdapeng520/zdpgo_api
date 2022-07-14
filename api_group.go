package zdpgo_api

import (
	"fmt"
	"strings"
	"unsafe"
)

// ApiGroup API分组
type ApiGroup struct {
	RouterGroup
}

func (g *ApiGroup) Any(method, routerPath string, handleFuncList ...func(ctx *Context)) {

	// 异常情况
	if handleFuncList == nil {
		return
	}

	// 处理结果的方法对象
	handleFuncObj := func(ctxGin *Context) {
		for _, handleFunc := range handleFuncList {
			ctxApiPointer := unsafe.Pointer(ctxGin) // 将父类实例转为通用指针
			ctxApi := (*Context)(ctxApiPointer)     //再转换为子类结构体
			handleFunc(ctxApi)
		}
	}

	switch strings.ToUpper(method) {
	case "GET":
		g.GET(routerPath, func(ctxGin *Context) {
			handleFuncObj(ctxGin)
		})
	case "POST":
		g.POST(routerPath, func(ctxGin *Context) {
			handleFuncObj(ctxGin)
		})
	case "PUT":
		g.PUT(routerPath, func(ctxGin *Context) {
			handleFuncObj(ctxGin)
		})
	case "DELETE":
		g.DELETE(routerPath, func(ctxGin *Context) {
			handleFuncObj(ctxGin)
		})
	case "PATCH":
		g.PATCH(routerPath, func(ctxGin *Context) {
			handleFuncObj(ctxGin)
		})
	case "HEAD":
		g.HEAD(routerPath, func(ctxGin *Context) {
			handleFuncObj(ctxGin)
		})
	default:
		fmt.Println("暂不支持此类型的方法", "method", method)
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
