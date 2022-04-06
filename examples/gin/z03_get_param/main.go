package main

import (
	"net/http"

	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
)

func main() {
	router := gin.Default()

	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("first_name", "张大鹏")
		lastname := c.Query("last_name") // 是 c.Request.URL.Query().Get("lastname") 的简写

		c.String(http.StatusOK, "你好 %s %s", firstname, lastname)
	})
	router.Run(":8080")
}
