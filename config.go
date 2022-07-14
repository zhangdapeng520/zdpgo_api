package zdpgo_api

type Config struct {
	Debug          bool             `yaml:"debug" json:"debug"`
	Host           string           `yaml:"host" json:"host"`                         // 启动地址，默认0.0.0.0
	Port           int              `yaml:"port" json:"port"`                         // 启动端口号，默认3333
	UploadFileSize int64            `yaml:"upload_file_size" json:"upload_file_size"` // 上传文件大小限制（M），默认33
	RateLimit      uint64           `yaml:"rate_limit" json:"rate_limit"`             // 请求速率，默认3333
	Middleware     MiddlewareConfig `yaml:"middleware" json:"middleware"`             // 中间件
	Router         RouterConfig     `yaml:"router" json:"router"`                     // 路由配置
}

type RouterConfig struct {
	HealthCheck bool `yaml:"health_check" json:"health_check"` // 健康检查路由
	Static      bool `yaml:"static" json:"static"`             // 静态文件路由
	Upload      bool `yaml:"upload" json:"upload"`             // 文件上传路由
}

type MiddlewareConfig struct {
	Cors      bool `yaml:"cors" json:"cors"`
	RateLimit bool `yaml:"rateLimit" json:"rateLimit"`
}
