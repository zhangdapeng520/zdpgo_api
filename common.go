package zdpgo_api

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_password"
)

/*
@Time : 2022/5/17 17:43
@Author : 张大鹏
@File : common.go
@Software: Goland2021.3.1
@Description: common 通用
*/

const (
	AuthUserKey = "user" // 是基本身份验证中用户凭据的cookie名称。
)

var (
	Password *zdpgo_password.Password
	Log      *zdpgo_log.Log
)

// JsonMap json字典类型
type JsonMap map[string]interface{}

// StringMap 字符串类型的字典
type StringMap map[string]string
