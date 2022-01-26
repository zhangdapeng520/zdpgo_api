package zdpgo_gin

import (
	"regexp"

	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func (g *Gin) initTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		g.trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, g.trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, g.trans)
		default:
			en_translations.RegisterDefaultTranslations(v, g.trans)
		}
		return
	}
	return
}

// validateMobile 手机号验证器
func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	return ok
}

// RegisterValidate 注册验证器
// @param key 校验关键字
// @param errorInfo 校验错误提示
// @param f 校验方法
func (g *Gin) RegisterValidate(key string, errorInfo string, f validator.Func) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 手机号验证器
		_ = v.RegisterValidation(key, f)

		// 手机号验证器翻译
		_ = v.RegisterTranslation(key, g.trans, func(ut ut.Translator) error {
			info := fmt.Sprintf("{0} %s", errorInfo)
			return ut.Add(key, info, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(key, fe.Field())
			return t
		})
	}
}

// registerValidates 注册内置的校验器
func (g *Gin) registerValidates() {
	g.RegisterValidate("mobile", "非法的手机号码", validateMobile)
}

func (g *Gin) registerValidateMobile1() {
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 手机号验证器
		_ = v.RegisterValidation("mobile", validateMobile)

		// 手机号验证器翻译
		_ = v.RegisterTranslation("mobile", g.trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}
