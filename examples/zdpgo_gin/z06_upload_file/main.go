package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的路径
		dst := "uploads/" + file.Filename
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println("上传文件失败：", err.Error())
		}

		// 响应
		c.JSON(200, gin.H{
			"status":    true,
			"file_name": file.Filename,
		})
	})
	g.Run()
}
