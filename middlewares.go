package zdpgo_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_code"
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
