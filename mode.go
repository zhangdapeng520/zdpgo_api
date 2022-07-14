package zdpgo_api

import (
	"io"
	"os"

	"github.com/zhangdapeng520/zdpgo_api/binding"
)

// EnvGinMode GIN模式的环境变量名称
const EnvGinMode = "GIN_MODE"

const (
	DebugMode   = "debug"   // debug模式
	ReleaseMode = "release" // 发布模式
	TestMode    = "test"    // 测试模式
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter 默认输出流
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter 默认错误输出流
var DefaultErrorWriter io.Writer = os.Stderr

var ginMode = debugCode
var modeName = DebugMode

// IsDebugging 是否为Debug模式
func IsDebugging() bool {
	return ginMode == debugCode
}

// 初始化方法
func init() {
	mode := os.Getenv(EnvGinMode) // 从环境变量中获取gin启动模式
	SetMode(mode)                 // 设置gin的模式
}

// SetMode 设置gin的模式
func SetMode(value string) {
	// 默认使用debug模式
	if value == "" {
		value = DebugMode
	}

	// 设置不同的模式
	switch value {
	case DebugMode:
		ginMode = debugCode
	case ReleaseMode:
		ginMode = releaseCode
	case TestMode:
		ginMode = testCode
	default:
		panic("错误的模式: " + value + " (仅支持: debug/release/est)")
	}

	// 返回模式名称
	modeName = value
}

// DisableBindValidation 关闭默认的校验器
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber 为绑定设置true。允许DecodeRusEnumber在JSON解码器实例上调用UseNumber方法。
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields 为绑定设置true。使EnableDecoderDisallowUnknowFields在JSON解码器实例上调用DisallowUnknowFields方法。
func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode 返回当前的gin模式
func Mode() string {
	return modeName
}
