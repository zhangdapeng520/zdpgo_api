package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zdpgo_gin/examples/z17_user/schemas"
)

// 发送邮件
func sendEmail(c *gin.Context) {

	// 自动提取邮箱信息
	email := schemas.SendEmail{}
	c.BindJSON(&email)
	log.Printf("%v", &email)

	// 数据校验

	// 数据入库

	// 直接返回
	c.JSON(http.StatusOK, email)
}
