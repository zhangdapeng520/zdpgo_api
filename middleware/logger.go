package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_zap"
	"go.uber.org/zap"
	"time"
)

// Logger 接收gin框架默认的日志
// @params records 要记录的内容。主要包括url，header，client，query，body
func Logger(log *zdpgo_zap.Zap, records []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)

		// 默认记录方式
		if records == nil {
			log.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
			return
		}

		// 根据不同的配置记录
		for _, record := range records {
			// 记录url日志
			// sScheme      string
			// Opaque      string    // encoded opaque data
			// User        *Userinfo // username and password information
			// Host        string    // host or host:port
			// Path        string    // path (relative paths may omit leading slash)
			// RawPath     string    // encoded path hint (see EscapedPath method)
			// ForceQuery  bool      // append a query ('?') even if RawQuery is empty
			// RawQuery    string    // encoded query values, without '?'
			// Fragment    string    // fragment for references, without '#'
			// RawFragment string    // encoded fragment hint (see EscapedFragment method)
			if record == "url" {
				log.Info("url相关信息",
					zap.String("Scheme", c.Request.URL.Scheme),
					zap.String("Opaque", c.Request.URL.Opaque),
					zap.Any("User", c.Request.URL.User),
					zap.String("Host", c.Request.URL.Host),
					zap.String("Path", c.Request.URL.Path),
					zap.String("RawPath", c.Request.URL.RawPath),
					zap.Bool("ForceQuery", c.Request.URL.ForceQuery),
					zap.String("RawQuery", c.Request.URL.RawQuery),
					zap.String("Fragment", c.Request.URL.Fragment),
					zap.String("RawFragment", c.Request.URL.RawFragment),
				)

			}

			// 记录header日志
			if record == "header" {
				//type Header map[string][]string
				log.Info("header相关信息", zap.Any("header", c.Request.Header))
			}

			// 记录form日志
			if record == "form" {
				//type Values map[string][]string
				log.Info("form相关信息", zap.Any("form", c.Request.Form), zap.Any("post_form", c.Request.PostForm))
			}

			// 记录body日志
			if record == "body" {
				//type Values map[string][]string
				log.Info("body相关信息", zap.Any("body", c.Request.Body))
			}
		}
	}
}
