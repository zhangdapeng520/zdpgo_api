package v1

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"log"
	"net/http"
	"zdpgo_gin/examples/z17_user/schemas"
)

// 根据用户名和密码登录
func loginUsername(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserLoginUsername{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}

// 根据邮箱和密码登录
func loginEmail(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserLoginEmail{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}

// 根据手机号和密码登录
func loginPhone(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserLoginPhone{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}

// 根据邮箱和验证码登录
func loginEmailCode(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserLoginEmailCode{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}

// 根据手机号和验证码登录
func loginPhoneCode(c *gin.Context) {

	// 自动提取注册信息
	user := schemas.UserLoginPhoneCode{}
	c.BindJSON(&user)
	log.Printf("%v", &user)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, user)
}
