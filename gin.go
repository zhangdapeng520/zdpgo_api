package zdpgo_api

import "github.com/zhangdapeng520/zdpgo_api/gin"

// NewGinWithLog 根据是否使用详细日志创建Gin的实例
// @param isDetailLogger 是否使用详细日志，开启后，会记录请求头，form，body的详细信息
func NewGinWithLog(isDetailLogger bool) *gin.Engine {
	// 创建gin对象
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		IsDetailLogger: isDetailLogger, // 使用详细日志
	}), gin.Recovery())

	// 设置模式
	if isDetailLogger {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 返回实例
	return r
}
