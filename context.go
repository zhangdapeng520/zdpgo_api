package zdpgo_api

import (
	"github.com/zhangdapeng520/zdpgo_api/gin"
)

/*
@Time : 2022/5/17 15:50
@Author : 张大鹏
@File : context.go
@Software: Goland2021.3.1
@Description: context上下文对象相关
*/

// Context 重命名gin上下文
type Context struct {
	gin.Context
}
