package routes

import "github.com/zhangdapeng520/zdpgo_api/libs/gin"

// Run 启动服务
func Run() {
	router := GetRoutes()
	router.Run(":5000")
}

// GetRoutes 获取路由
func GetRoutes() *gin.Engine {
	// 创建服务
	router := gin.Default()

	// 路由分组1
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	// 路由分组2
	v2 := router.Group("/v2")
	addPingRoutes(v2)

	// 返回服务对象
	return router
}
