# zdpgo_gin
基于gin二次封装的一个后端api快速开发框架

项目地址：https://github.com/zhangdapeng520/zdpgo_gin

## 功能清单
- 根据MySQL数据库表自动生成REST风格的CRUD API接口

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

## 示例
### 根据MySQL数据库表自动生成REST风格的CRUD API接口
会自动生成以下接口：
- 根据ID查询数据
- 新增数据
- 批量新增数据
- 根据ID修改数据
- 根据ID列表修改数据
- 根据ID删除数据
- 根据ID列表删除数据
- 根据ID列表查询数据
- 根据分页查询数据

```shell
package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_mysql"
)

type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 设置MySQL
	g.SetMysql(zdpgo_mysql.MysqlConfig{
		Debug:    true,
		Host:     "192.168.33.101",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	})

	// 创建app
	app := gin.Default()

	// 创建路由组
	group := app.Group("/api/v1")

	// 注册路由
	var students []Student
	g.RegisterCrudRouter(group, "student", &students)

	// 启动app
	app.Run("0.0.0.0:8888")
}
```
