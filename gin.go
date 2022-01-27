package zdpgo_gin

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/zhangdapeng520/zdpgo_zap"
)

// Gin 核心对象
type Gin struct {
	log    *zdpgo_zap.Zap // 日志对象
	config *GinConfig     // 配置对象
	trans  ut.Translator  // 翻译器对象
	err    error          // 错误信息对象，记录最近的错误
	app    *gin.Engine    // app核心对象
}

// GinConfig 配置对象
type GinConfig struct {
	Debug       bool   // 是否为debug模式
	LogFilePath string // 日志路径
	Language    string // 语言，默认zh中文
	JwtKey      string // jwt权限校验的key
	JwtExpired  int64  // jwt过期时间（秒）
}

// New 生成Gin对象
func New(config GinConfig) *Gin {
	g := Gin{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_gin.log"
	}
	l := zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   true,
		OpenFileName: true,
		LogFilePath:  config.LogFilePath,
	})

	g.log = l

	// 初始化翻译
	if config.Language == "" {
		config.Language = "zh"
	}
	g.err = g.initTrans(config.Language)
	if g.err != nil {
		g.log.Error("初始化翻译器失败", "error", g.err.Error())
	}

	// 初始化jwt
	if config.JwtKey == "" {
		config.JwtKey = "zhangdapengZHANGDAPENG!@#$%^&*()_+123456789"
	}
	if config.JwtExpired == 0 {
		config.JwtExpired = 60 * 60 * 3 // 3小时
	}

	// 注册内置的校验器
	g.registerValidates()

	// 初始化配置
	g.config = &config

	return &g
}
