package zdpgo_api

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"unsafe"
)

/*
@Time : 2022/5/16 20:29
@Author : 张大鹏
@File : basic_auth.go
@Software: Goland2021.3.1
@Description: basic_auth基本权限相关
*/

// GetBasicAuthGroup 获取基本权限校验路由组
func (a *Api) GetBasicAuthGroup(routerPath string, accounts StringMap) (group *ApiGroup) {
	// 将api的账户转换为gin的账户
	accountsPointer := unsafe.Pointer(&accounts)
	mapString := (*map[string]string)(accountsPointer)

	// 获取gin的分组
	ginGroup := a.App.Group(routerPath, gin.BasicAuth(*mapString))

	// 转换为api的分组
	ginGroupPointer := unsafe.Pointer(ginGroup)
	group = (*ApiGroup)(ginGroupPointer)

	return
}
