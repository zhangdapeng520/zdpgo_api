package zdpgo_api

import (
	"encoding/base64"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"
)

// RouterHealthCheck 健康检查处理器
func RouterHealthCheck(ctx *Context) {
	success := ctx.GetResponseSuccess(nil)
	ctx.JSON(200, success)
}

// AddHealthCheckRouter 添加健康检查接口
func (a *Api) AddHealthCheckRouter() {
	a.Get("/health", RouterHealthCheck)
}

// AddUploadRouter 添加文件上传路由
// @param routerPath 路由路径，比如：/upload
// @param fileName 上传文件的表单名称，比如：file
// @param saveDir 文件的上传路径，比如：./uploads
// @param handleResultList 处理结果的方法列表，用户可以自定义文件上传后的返回内容，处理错误信息
func (a *Api) AddUploadRouter(
	routerPath string,
	fileName string,
	saveDir string,
	handleResultList ...func(c *Context, file *multipart.FileHeader, err error)) {

	// 处理结果的方法
	handleResultFunc := func(c *Context, file *multipart.FileHeader, err error) {
		if handleResultList != nil {
			for _, handleResult := range handleResultList {
				handleResult(c, file, err)
			}
		} else { // 默认的文件响应处理方式
			// 准备响应对象
			response := c.GetResponseSuccess(nil)
			data := make(map[string]string)

			// 获取要保存的文件名
			filename := path.Join(saveDir, filepath.Base(file.Filename))
			data["filename"] = file.Filename

			// 文件内容
			fileContent, err := ioutil.ReadFile(filename)
			if err != nil {
				data["file_content"] = base64.StdEncoding.EncodeToString([]byte(""))
			} else {
				data["file_content"] = base64.StdEncoding.EncodeToString(fileContent)
			}

			// 计算MD5值
			data["md5"] = GetMd5(string(fileContent))
			response.Data = data
			c.JSON(200, response)
		}
	}

	// 添加POST类型的路由
	a.Post(routerPath, func(c *Context) {
		// 获取上传的文件
		file, err := c.FormFile(fileName)
		if err != nil {
			handleResultFunc(c, file, err)
			return
		}

		// 创建要保存的文件夹
		err = CreateDirs(saveDir)
		if err != nil {
			handleResultFunc(c, file, err)
			return
		}

		// 获取要保存的文件名
		filename := path.Join(saveDir, filepath.Base(file.Filename))

		// 保存上传文件
		if err = c.SaveUploadedFile(file, filename); err != nil {
			handleResultFunc(c, file, err)
			return
		}

		// 处理返回结果
		handleResultFunc(c, file, err)
	})
}

// AddStaticRouter 添加静态文件目录
func (a *Api) AddStaticRouter(routerPath string, dirPath string) {
	a.App.StaticFS(routerPath, http.Dir(dirPath))
}

// AddStaticFSRouter 添加嵌入式文件系统作为静态目录
func (a *Api) AddStaticFSRouter(routerPath string, fsObj fs.FS) {
	a.App.StaticFS(routerPath, http.FS(fsObj))
}
