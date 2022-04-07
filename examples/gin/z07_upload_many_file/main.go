package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
	"log"
)

func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/test", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		fmt.Println("form：", form)
		files := form.File["file"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			dst := "uploads/" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(200, gin.H{
			"status": true,
		})
	})
	router.Run(":8080")
}
