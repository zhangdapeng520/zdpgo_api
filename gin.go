package zdpgo_gin

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/zhangdapeng520/zdpgo_mysql"
	"github.com/zhangdapeng520/zdpgo_zap"
	"net/http"
)

// Gin 核心对象
type Gin struct {
	log    *zdpgo_zap.Zap     // 日志对象
	config *GinConfig         // 配置对象
	trans  ut.Translator      // 翻译器对象
	err    error              // 错误信息对象，记录最近的错误
	App    *gin.Engine        // app核心对象，可以被外部访问，用于加载路由
	mysql  *zdpgo_mysql.Mysql // mysql核心对象
}

// New 生成Gin对象
func New(config GinConfig) *Gin {
	g := Gin{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_gin.log"
	}
	l := zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   false,
		OpenFileName: true,
		LogFilePath:  config.LogFilePath,
	})

	g.log = l

	// 初始化app
	g.App = gin.New()                 // 创建app
	g.App.Use(g.MiddlewareCors())     // 使用跨域中间件
	g.App.Use(g.MiddlewareLogger())   // 使用日志中间件
	g.App.Use(g.MiddlewareRecovery()) // 使用异常捕获中间件

	// 注册通用路由
	if config.OpenCommonRouter {
		g.RegisterCommonRouter(g.App)
	}
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 加载模板
	if config.TemplatePath != "" {
		g.log.Info("加载模板", "templatePath", config.TemplatePath)
		g.App.LoadHTMLGlob(config.TemplatePath) // 加载模板
	}

	// 加载静态目录
	if config.StaticPath != "" {
		if config.StaticUrl == "" {
			config.StaticUrl = "/static"
		}
		g.log.Info("指定静态目录", "url", config.StaticUrl, "path", config.StaticPath)
		g.App.StaticFS(config.StaticUrl, http.Dir(config.StaticPath)) // 指定静态目录
	}

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

// SetMysql 设置MySQL
func (g *Gin) SetMysql(config zdpgo_mysql.MysqlConfig) {
	if g.mysql == nil {
		g.mysql = zdpgo_mysql.New(config)
	}
}
