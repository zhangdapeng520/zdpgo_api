package zdpgo_gin

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/zhangdapeng520/zdpgo_code"
)

// 默认内存存储的方式
var store = base64Captcha.DefaultMemStore

// GetRouterCommonCaptcha 获取图片验证码
func (g *Gin) GetRouterCommonCaptcha(ctx *gin.Context) {
	// 创建驱动
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	// 创建验证码对象
	cp := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, err := cp.Generate()

	// 创建并返回响应
	rsp := NewResponse()
	if err != nil {
		g.log.Error("生成验证码错误", "error", err.Error())
		rsp.Code = zdpgo_code.CODE_PARAM_ERROR
		rsp.Message = "生成验证码错误"
		g.Success(ctx, rsp)
		return
	}

	// 创建并返回数据响应
	data := gin.H{
		"captchaId": id,
		"picPath":   b64s,
	}
	rspData := NewResponseData(data)
	g.SuccessData(ctx, rspData)
}

// CaptchaVerify 图片验证码校验
// @param captchaId 验证码id
// @param captcha 验证码
// @return 校验成功返回true，失败返回false
func (g *Gin) CaptchaVerify(captchaId, captcha string) bool {
	return store.Verify(captchaId, captcha, true)
}
