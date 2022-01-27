package zdpgo_gin

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zhangdapeng520/zdpgo_code"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func (g *Gin) HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	rsp := NewResponse()
	if !ok {
		rsp.Status = false
		rsp.Code = zdpgo_code.CODE_PARAM_ERROR
		rsp.Message = err.Error()
		g.Success(c, rsp)
		return
	}
	data := removeTopStruct(errs.Translate(g.trans))
	rspData := NewResponseData(data)
	rspData.Code = zdpgo_code.CODE_PARAM_ERROR
	rspData.Message = zdpgo_code.MESSAGE_PARAM_ERROR
	rspData.Status = false
	g.SuccessData(c, rspData)
	return
}
