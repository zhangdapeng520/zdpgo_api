package gin

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

// Log 核心日志对象
var Log *zdpgo_log.Log

func init() {
	if Log == nil {
		Log = zdpgo_log.NewWithDebug(true, "logs/zdpgo/zdpgo_api.log")
	}
}

// Logger 根据配置返回一个logger日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		// 通过请求
		c.Next()

		// 输出日志
		Log.Debug("ZDP-Go-Api调试日志",
			"method", c.Request.Method,
			"path", c.Request.URL.String(),
			"status_code", c.Writer.Status(),
			"header", c.Request.Header,
		)
	}
}
