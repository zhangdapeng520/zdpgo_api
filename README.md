# zdpgo_api
基于gin二次封装的一个后端api快速开发框架

项目地址：https://github.com/zhangdapeng520/zdpgo_api

## 版本历史
- 版本1.0.0：2022年2月9日
- 版本1.0.1：2022年2月11日 指定默认模板目录template和默认静态文件夹目录static
- 版本1.0.2：2022年2月11日 移除本地依赖
- 版本1.0.3：2022年2月11日 新增日志中间件和recover中间件
- 版本1.0.4：2022年3月8日 优化模板和静态目录
- 版本1.0.5：2022年3月8日 增加常用的统一返回对象
- 版本1.0.6：2022年3月8日 将挂载通用路由设置为开关量
- 版本1.0.7：2022年3月10日 支持配置服务相关参数
- 版本1.0.8：2022年3月10日 支持viper读取和设置配置
- 版本1.0.9：2022年3月11日 logger日志支持配置
- 版本1.1.0：2022年3月12日 修复日志不打印body
- 版本1.1.1：2022年4月8日 新增详细日志
- 版本1.1.2：2022年4月11日 详细日志升级，支持查看的具体Body中的JSON数据
- 版本1.1.3：2022年4月21日 文件上传

## 示例
### 文件上传
```go
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
```