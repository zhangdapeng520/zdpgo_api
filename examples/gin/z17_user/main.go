package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"zdpgo_gin/examples/z17_user/routers"
)

func main() {
	// 创建app
	app := gin.Default()

	// 注册路由
	routers.RegisterRouter(app)

	// 启动app
	app.Run(":8080")
}
