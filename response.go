package zdpgo_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_code"
)

// Response 响应
type Response struct {
	Code    uint32 `json:"code"`    // 状态码
	Message string `json:"message"` // 消息
	Status  bool   `json:"status"`  // 状态
}

// NewResponse 创建默认的响应对象
func NewResponse() Response {
	resp := Response{
		Code:    zdpgo_code.CODE_SUCCESS,
		Message: zdpgo_code.MESSAGE_SUCCESS,
		Status:  true,
	}
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
func NewResponseData(data interface{}) ResponseData {
	resp := ResponseData{
		Code:    zdpgo_code.CODE_SUCCESS,
		Message: zdpgo_code.MESSAGE_SUCCESS,
		Status:  true,
		Data:    data,
	}
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
