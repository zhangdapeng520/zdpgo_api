package zdpgo_gin

// GinConfig 配置对象
type GinConfig struct {
	Debug        bool   // 是否为debug模式
	LogFilePath  string // 日志路径
	Language     string // 语言，默认zh中文
	JwtKey       string // jwt权限校验的key
	JwtExpired   int64  // jwt过期时间（秒）
	TemplatePath string // 模板路径
	StaticPath   string // 静态文件路径
	StaticUrl    string // 静态文件路由
}
