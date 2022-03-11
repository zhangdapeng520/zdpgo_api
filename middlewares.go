package zdpgo_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_code"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// MiddlewareCors 跨域中间件
func (g *Gin) MiddlewareCors() gin.HandlerFunc {
	// 返回中间件函数
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token, zdp-token, ZDPToken")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

// MiddlewareJWTAuth jwt权限校验中间件
func (g *Gin) MiddlewareJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Request.Header.Get("ZDPtoken")
		g.log.Info("token信息", "token", token)
		rspData := NewResponseData(nil)

		// 处理token不存在
		if token == "" {
			rspData.Code = zdpgo_code.CODE_UN_AUTH
			rspData.Message = zdpgo_code.MESSAGE_UN_AUTH
			g.SuccessData(c, rspData)
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := g.ParseToken(token)
		if err != nil {
			g.log.Error("解析token失败", "error", err.Error())
			if err == TokenExpired {
				if err == TokenExpired {
					rsp := Response{
						Code:    zdpgo_code.CODE_TOKEN_EXPIRED,
						Message: zdpgo_code.MESSAGE_TOKEN_EXPIRED,
						Status:  false,
					}
					g.Success(c, rsp)
					c.Abort()
					return
				}
			}

			rsp := Response{
				Code:    zdpgo_code.CODE_UN_AUTH,
				Message: zdpgo_code.MESSAGE_UN_AUTH,
				Status:  false,
			}
			g.Success(c, rsp)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

// MiddlewareIsAdmin1Auth 一级管理员权限校验
func (g *Gin) MiddlewareIsAdmin1Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*Claims)
		g.log.Info("当前用户", "currentUser", currentUser, "role", currentUser.Role)

		if currentUser.Role != 1 {
			rsp := Response{
				Code:    zdpgo_code.CODE_UN_AUTH,
				Message: zdpgo_code.MESSAGE_UN_AUTH,
				Status:  false,
			}
			g.Success(ctx, rsp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}

// MiddlewareIsAdmin2Auth 二级管理员权限校验
func (g *Gin) MiddlewareIsAdmin2Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*Claims)

		if currentUser.Role != 1 && currentUser.Role != 2 {
			rsp := Response{
				Code:    zdpgo_code.CODE_UN_AUTH,
				Message: zdpgo_code.MESSAGE_UN_AUTH,
				Status:  false,
			}
			g.Success(ctx, rsp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}

// MiddlewareIsAdmin3Auth 三级管理员权限校验
func (g *Gin) MiddlewareIsAdmin3Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*Claims)

		if currentUser.Role != 1 && currentUser.Role != 2 && currentUser.Role != 3 {
			rsp := Response{
				Code:    zdpgo_code.CODE_UN_AUTH,
				Message: zdpgo_code.MESSAGE_UN_AUTH,
				Status:  false,
			}
			g.Success(ctx, rsp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}

// MiddlewareLogger 接收gin框架默认的日志
func (g *Gin) MiddlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		g.log.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)

		// 默认记录方式
		if g.config.Server.Records != nil {
			// 根据不同的配置记录
			for _, record := range g.config.Server.Records {
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
					g.log.Info("url相关信息",
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
					g.log.Info("header相关信息", zap.Any("header", c.Request.Header))
				}

				// 记录form日志
				if record == "form" {
					//type Values map[string][]string
					g.log.Info("form相关信息", zap.Any("form", c.Request.Form), zap.Any("post_form", c.Request.PostForm))
				}

				// 记录body日志
				if record == "body" {
					//type Values map[string][]string
					g.log.Info("body相关信息", zap.Any("body", c.Request.Body))
				}
			}
		}
	}
}

// MiddlewareRecovery recover掉项目可能出现的panic
func (g *Gin) MiddlewareRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					g.log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				g.log.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
