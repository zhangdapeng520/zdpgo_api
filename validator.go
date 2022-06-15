package zdpgo_api

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/zhangdapeng520/zdpgo_api/validator"
	zhs "github.com/zhangdapeng520/zdpgo_api/validator/translations/zh"
)

/*
@Time : 2022/5/16 19:56
@Author : 张大鹏
@File : validator.go
@Software: Goland2021.3.1
@Description: validator验证器相关方法
*/
var (
	validate = validator.New()          // 实例化验证器
	chinese  = zh.New()                 // 获取中文翻译器
	uni      = ut.New(chinese, chinese) // 设置成中文翻译器
	trans, _ = uni.GetTranslator("zh")  // 获取翻译字典
)

/**
常用的规则：
标记	标记说明	例
required	必填	Field或Struct validate:"required"
omitempty	空时忽略	Field或Struct validate:"omitempty"
len	长度	Field validate:"len=0"
eq	等于	Field validate:"eq=0"
gt	大于	Field validate:"gt=0"
gte	大于等于	Field validate:"gte=0"
lt	小于	Field validate:"lt=0"
lte	小于等于	Field validate:"lte=0"
eqfield	同一结构体字段相等	Field validate:"eqfield=Field2"
nefield	同一结构体字段不相等	Field validate:"nefield=Field2"
gtfield	大于同一结构体字段	Field validate:"gtfield=Field2"
gtefield	大于等于同一结构体字段	Field validate:"gtefield=Field2"
ltfield	小于同一结构体字段	Field validate:"ltfield=Field2"
ltefield	小于等于同一结构体字段	Field validate:"ltefield=Field2"
eqcsfield	跨不同结构体字段相等	Struct1.Field validate:"eqcsfield=Struct2.Field2"
necsfield	跨不同结构体字段不相等	Struct1.Field validate:"necsfield=Struct2.Field2"
gtcsfield	大于跨不同结构体字段	Struct1.Field validate:"gtcsfield=Struct2.Field2"
gtecsfield	大于等于跨不同结构体字段	Struct1.Field validate:"gtecsfield=Struct2.Field2"
ltcsfield	小于跨不同结构体字段	Struct1.Field validate:"ltcsfield=Struct2.Field2"
ltecsfield	小于等于跨不同结构体字段	Struct1.Field validate:"ltecsfield=Struct2.Field2"
min	最大值	Field validate:"min=1"
max	最小值	Field validate:"max=2"
structonly	仅验证结构体，不验证任何结构体字段	Struct validate:"structonly"
nostructlevel	不运行任何结构级别的验证	Struct validate:"nostructlevel"
dive	向下延伸验证，多层向下需要多个dive标记	[][]string validate:"gt=0,dive,len=1,dive,required"
dive Keys & EndKeys	与dive同时使用，用于对map对象的键的和值的验证，keys为键，endkeys为值	map[string]string `validate:”gt=0,dive,keys,eq=1
required_with	其他字段其中一个不为空且当前字段不为空	Field validate:"required_with=Field1 Field2"
required_with_all	其他所有字段不为空且当前字段不为空	Field validate:"required_with_all=Field1 Field2"
required_without	其他字段其中一个为空且当前字段不为空	Field `validate:”required_without=Field1 Field2”
required_without_all	其他所有字段为空且当前字段不为空	Field validate:"required_without_all=Field1 Field2"
isdefault	是默认值	Field validate:"isdefault=0"
oneof	其中之一	Field validate:"oneof=5 7 9"
containsfield	字段包含另一个字段	Field validate:"containsfield=Field2"
excludesfield	字段不包含另一个字段	Field validate:"excludesfield=Field2"
unique	是否唯一，通常用于切片或结构体	Field validate:"unique"
alphanum	字符串值是否只包含 ASCII 字母数字字符	Field validate:"alphanum"
alphaunicode	字符串值是否只包含 unicode 字符	Field validate:"alphaunicode"
alphanumunicode	字符串值是否只包含 unicode 字母数字字符	Field validate:"alphanumunicode"
numeric	字符串值是否包含基本的数值	Field validate:"numeric"
hexadecimal	字符串值是否包含有效的十六进制	Field validate:"hexadecimal"
hexcolor	字符串值是否包含有效的十六进制颜色	Field validate:"hexcolor"
lowercase	符串值是否只包含小写字符	Field validate:"lowercase"
uppercase	符串值是否只包含大写字符	Field validate:"uppercase"
email	字符串值包含一个有效的电子邮件	Field validate:"email"
json	字符串值是否为有效的 JSON	Field validate:"json"
file	符串值是否包含有效的文件路径，以及该文件是否存在于计算机上	Field validate:"file"
url	符串值是否包含有效的 url	Field validate:"url"
uri	符串值是否包含有效的 uri	Field validate:"uri"
base64	字符串值是否包含有效的 base64值	Field validate:"base64"
contains	字符串值包含子字符串值	Field validate:"contains=@"
containsany	字符串值包含子字符串值中的任何字符	Field validate:"containsany=abc"
containsrune	字符串值包含提供的特殊符号值	Field validate:"containsrune=☢"
excludes	字符串值不包含子字符串值	Field validate:"excludes=@"
excludesall	字符串值不包含任何子字符串值	Field validate:"excludesall=abc"
excludesrune	字符串值不包含提供的特殊符号值	Field validate:"containsrune=☢"
startswith	字符串以提供的字符串值开始	Field validate:"startswith=abc"
endswith	字符串以提供的字符串值结束	Field validate:"endswith=abc"
ip	字符串值是否包含有效的 IP 地址	Field validate:"ip"
ipv4	字符串值是否包含有效的 ipv4地址	Field validate:"ipv4"
datetime	字符串值是否包含有效的 日期	Field validate:"datetime"
*/

// InitValidator 初始化校验器
func (a *Api) InitValidator() (err error) {
	// 注册翻译器
	err = zhs.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		Log.Error("初始化校验器失败", "error", err)
		return
	}

	return
}

// Validate 校验结构体
func (a *Api) Validate(data interface{}) (errData map[string]string) {
	err := validate.Struct(data)
	if err != nil {
		Log.Error("校验数据失败", "error", err, "data", data)
		if errors, ok := err.(validator.ValidationErrors); ok {
			errData = errors.Translate(trans)
			return
		}
	}
	return
}
