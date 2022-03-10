package zdpgo_gin

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`         // redis主机地址
	Port     uint16 `mapstructure:"port" json:"port"`         // redis端口号
	Database uint   `mapstructure:"database" json:"database"` // redis数据库
	Size     uint   `mapstructure:"size" json:"size"`         // redis最大连接数
}

// SessionConfig session配置
type SessionConfig struct {
	OpenSession bool        `mapstructure:"open_session" json:"open_session"` // 使用启用session
	Key         string      `mapstructure:"key" json:"key"`                   // 加密秘钥
	SessionName string      `mapstructure:"session_name" json:"session_name"` // session名字
	SessionType string      `mapstructure:"session_type" json:"session_type"` // session类型：cookie或redis
	Redis       RedisConfig `mapstructure:"redis" json:"redis"`               // redis相关配置
}

// JwtConfig jwt配置
type JwtConfig struct {
	Key     string `mapstructure:"key" json:"key"`         // jwt权限校验的key
	Expired int64  `mapstructure:"expired" json:"expired"` // jwt过期时间（秒）
}

// ServerConfig 服务配置
type ServerConfig struct {
	Host           string `mapstructure:"host" json:"host"`                         // 服务启动的主机地址
	Port           uint16 `mapstructure:"port" json:"port"`                         // 服务启动的端口号
	ReadTimeout    uint16 `mapstructure:"read_timeout" json:"read_timeout"`         // 读超时时间
	WriteTimeout   uint16 `mapstructure:"write_timeout" json:"write_timeout"`       // 写超时时间
	MaxHeaderBytes uint32 `mapstructure:"max_header_bytes" json:"max_header_bytes"` // 请求头大小限制
}

// GinConfig 配置对象
type GinConfig struct {
	Debug            bool          `mapstructure:"debug" json:"debug"`                           // 是否为debug模式
	LogFilePath      string        `mapstructure:"log_file_path" json:"log_file_path"`           // 日志路径
	Language         string        `mapstructure:"language" json:"language"`                     // 语言，默认zh中文
	TemplatePath     string        `mapstructure:"template_path" json:"template_path"`           // 模板路径
	StaticPath       string        `mapstructure:"static_path" json:"static_path"`               // 静态文件路径
	StaticUrl        string        `mapstructure:"static_url" json:"static_url"`                 // 静态文件路由
	OpenCommonRouter bool          `mapstructure:"open_common_router" json:"open_common_router"` // 是否挂载通用路由
	OpenWebsocket    bool          `mapstructure:"open_websocket" json:"open_websocket"`         // 是否开启websocket通信
	Session          SessionConfig `mapstructure:"session" json:"session"`                       // session配置
	Jwt              JwtConfig     `mapstructure:"jwt" json:"jwt"`                               // jwt配置
	Server           ServerConfig  `mapstructure:"server" json:"server"`                         // jwt配置
}
