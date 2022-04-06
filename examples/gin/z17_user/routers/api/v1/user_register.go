package v1

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"log"
	"net/http"
	"zdpgo_gin/examples/z17_user/schemas"
)

// 根据用户名和密码注册
func registerUsername(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserRegisterUsername{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}

// 根据邮箱和密码注册
func registerEmail(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserRegisterEmail{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}

// 根据手机号和密码注册
func registerPhone(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserRegisterPhone{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}
