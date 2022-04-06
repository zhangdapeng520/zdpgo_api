package main

import (
	"github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
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
