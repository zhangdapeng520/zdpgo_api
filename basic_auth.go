package zdpgo_api

import "github.com/zhangdapeng520/zdpgo_api/gin"

/*
@Time : 2022/5/16 20:29
@Author : 张大鹏
@File : basic_auth.go
@Software: Goland2021.3.1
@Description: basic_auth基本权限相关
*/

// GetBasicAuthGroup 获取基本权限校验路由组
func (a *Api) GetBasicAuthGroup(routerPath string, accounts map[string]string) (group *gin.RouterGroup) {
	group = a.App.Group(routerPath, gin.BasicAuth(accounts))
	return
}
