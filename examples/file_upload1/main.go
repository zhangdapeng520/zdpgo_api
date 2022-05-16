package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_api/core/router"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"mime/multipart"
	"net/http"
)

func main() {
	app := gin.Default()
	router.Upload(app, 8, "/upload", "file", "./uploads", func(c *gin.Context, file *multipart.FileHeader, err error) {
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("文件上传失败: %s", err.Error()))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("文件 %s 上传成功 ", file.Filename))
	})
	app.Run(":8888")
}
