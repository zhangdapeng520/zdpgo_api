package zdpgo_gin

// SessionConfig session配置
type SessionConfig struct {
	OpenSession bool   // 使用启用session
	Key         string // 加密秘钥
	SessionName string // session名字
	SessionType string // session类型：cookie或redis
	RedisHost   string // redis主机地址
	RedisPort   uint16 // redis端口号
	RedisDb     uint   // redis数据库
	RedisSize   uint   // redis最大连接数
}

// JwtConfig jwt配置
type JwtConfig struct {
	JwtKey     string // jwt权限校验的key
	JwtExpired int64  // jwt过期时间（秒）
}

// GinConfig 配置对象
type GinConfig struct {
	Debug            bool          // 是否为debug模式
	LogFilePath      string        // 日志路径
	Language         string        // 语言，默认zh中文
	TemplatePath     string        // 模板路径
	StaticPath       string        // 静态文件路径
	StaticUrl        string        // 静态文件路由
	OpenCommonRouter bool          // 是否挂载通用路由
	OpenWebsocket    bool          // 是否开启websocket通信
	Session          SessionConfig // session配置
	Jwt              JwtConfig     // jwt配置
}
