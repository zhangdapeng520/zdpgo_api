package v1

import "github.com/gin-gonic/gin"

var (
	v1 *gin.RouterGroup
)

// 注册V1版本接口
func RegisterV1(app *gin.Engine) {
	v1 = app.Group("/api/v1")
	v1.POST("/register/username", registerUsername) // 注册
	v1.POST("/login/username", loginUsername)       // 登录
	v1.POST("/register/email", registerEmail)       // 邮箱注册
	v1.POST("/login/email", loginEmail)             // 邮箱登录
	v1.POST("/register/phone", registerPhone)       // 手机号注册
	v1.POST("/login/phone", loginPhone)             // 手机号登录
	v1.POST("/login/phone/code", loginPhoneCode)    // 手机号和验证码登录
	v1.POST("/login/email/code", loginEmailCode)    // 邮箱和验证码登录
	v1.POST("/send_email", sendEmail)               // 发送邮件

}
