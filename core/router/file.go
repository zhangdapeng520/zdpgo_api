package router

import (
	"github.com/zhangdapeng520/zdpgo_api/core/util"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"mime/multipart"
	"path"
	"path/filepath"
)

// Upload 文件上传路由
// @param fileSizeM 限制文件大小（M）
// @param routerPath 路由路径，比如：/upload
// @param fileName 上传文件的表单名称，比如：file
// @param saveDir 文件的上传路径，比如：./uploads
// @param handleResult 处理结果的方法，用户可以自定义文件上传后的返回内容，处理错误信息
func Upload(router *gin.Engine, fileSizeM int64, routerPath string, fileName string,
	saveDir string, handleResult func(c *gin.Context, file *multipart.FileHeader, err error)) {
	router.MaxMultipartMemory = fileSizeM << 20 // 8 MiB
	router.POST(routerPath, func(c *gin.Context) {
		file, err := c.FormFile(fileName)
		if err != nil {
			handleResult(c, nil, err)
			return
		}

		err = util.MakeDirs(saveDir)
		if err != nil {
			handleResult(c, nil, err)
			return
		}

		filename := path.Join(saveDir, filepath.Base(file.Filename))
		if err := c.SaveUploadedFile(file, filename); err != nil {
			handleResult(c, nil, err)
			return
		}

		// 处理返回结果
		handleResult(c, file, err)
	})
}
