package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
	"log"
)

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	g.App.POST("/test", func(c *gin.Context) {
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
	g.Run()
}
