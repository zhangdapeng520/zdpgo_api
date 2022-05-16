package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_api/core/util"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		basePath := "./uploads/"
		err = util.MakeDirs(basePath)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("create dir err: %s", err.Error()))
			return
		}

		filename := basePath + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("uploads file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("文件 %s 上传成功 ", file.Filename))
	})
	router.Run(":8888")
}
