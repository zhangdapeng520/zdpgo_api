package zdpgo_gin

import (
	"net/http"

	"github.com/zhangdapeng520/zdpgo_code"
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
)

// Response 响应
type Response struct {
	Code    uint32 `json:"code"`    // 状态码
	Message string `json:"message"` // 消息
	Status  bool   `json:"status"`  // 状态
}

// NewResponse 创建默认的响应对象
func NewResponse(messages ...string) Response {
	// 消息描述
	var msg string = zdpgo_code.MESSAGE_SUCCESS
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 响应
	resp := Response{
		Code:    zdpgo_code.CODE_SUCCESS,
		Message: msg,
		Status:  true,
	}

	// 返回响应
	return resp
}

// ResponseData 带数据的响应
type ResponseData struct {
	Code    uint32      `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息
	Status  bool        `json:"status"`  // 状态
	Data    interface{} `json:"data"`    // 数据
}

// NewResponseData 创建响应数据
func NewResponseData(data interface{}, messages ...string) ResponseData {
	// 消息描述
	var msg string = zdpgo_code.MESSAGE_SUCCESS
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 响应
	resp := ResponseData{
		Code:    zdpgo_code.CODE_SUCCESS,
		Message: msg,
		Status:  true,
		Data:    data,
	}

	// 返回响应
	return resp
}

// Success 返回成功的响应
func (g *Gin) Success(ctx *gin.Context, resp Response) {
	ctx.JSON(http.StatusOK, resp)
}

// SuccessData 返回成功的数据响应
func (g *Gin) SuccessData(ctx *gin.Context, resp ResponseData) {
	ctx.JSON(http.StatusOK, resp)
}

// ParamError 参数错误
func (g *Gin) ParamError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_PARAM_ERROR
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_PARAM_ERROR,
		Message: msg,
		Status:  false,
	})
}

// ServerError 服务器内部错误
func (g *Gin) ServerError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_SERVER_ERROR
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_SERVER_ERROR,
		Message: msg,
		Status:  false,
	})
}

// NotFoundError 资源不存在
func (g *Gin) NotFoundError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_NOT_FOUND
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_NOT_FOUND,
		Message: msg,
		Status:  false,
	})
}

// GrpcCanNotUseError gRPC服务不可用
func (g *Gin) GrpcCanNotUseError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_GRPC_NOT_USE
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_GRPC_NOT_USE,
		Message: msg,
		Status:  false,
	})
}

// ExistsError 数据已存在
func (g *Gin) ExistsError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_EXISTS_ERROR
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_EXISTS_ERROR,
		Message: msg,
		Status:  false,
	})
}

// UnAuthError 没有访问权限
func (g *Gin) UnAuthError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_UN_AUTH
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_UN_AUTH,
		Message: msg,
		Status:  false,
	})
}

// TokenExpiredError token已过期
func (g *Gin) TokenExpiredError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_TOKEN_EXPIRED
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_TOKEN_EXPIRED,
		Message: msg,
		Status:  false,
	})
}

// TimeOutError 连接超时
func (g *Gin) TimeOutError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_TIMEOUT
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_TIMEOUT,
		Message: msg,
		Status:  false,
	})
}

// CorsError 跨域请求失败
func (g *Gin) CorsError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_CORS_ERROR
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_CORS_ERROR,
		Message: msg,
		Status:  false,
	})
}

// RequestLimitError 请求限制
func (g *Gin) RequestLimitError(ctx *gin.Context, messages ...string) {
	// 错误提示
	var msg string = zdpgo_code.MESSAGE_REQUEST_LIMIT_ERROR
	if len(messages) > 0 {
		msg = messages[0]
	}

	// 返回json数据
	ctx.JSON(http.StatusOK, Response{
		Code:    zdpgo_code.CODE_REQUEST_LIMIT_ERROR,
		Message: msg,
		Status:  false,
	})
}
