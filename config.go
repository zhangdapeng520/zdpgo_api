package zdpgo_api

/*
@Time : 2022/5/16 16:25
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: config配置相关
*/

type Config struct {
	Debug          bool             `yaml:"debug" json:"debug"`
	Host           string           `yaml:"host" json:"host"`                         // 启动地址，默认0.0.0.0
	Port           int              `yaml:"port" json:"port"`                         // 启动端口号，默认3333
	UploadFileSize int64            `yaml:"upload_file_size" json:"upload_file_size"` // 上传文件大小限制（M），默认33
	Ecc            EccConfig        `yaml:"ecc" json:"ecc"`
	RateLimit      uint64           `yaml:"rate_limit" json:"rate_limit"` // 请求速率，默认3333
	Middleware     MiddlewareConfig `yaml:"middleware" json:"middleware"` // 中间件
	Router         RouterConfig     `yaml:"router" json:"router"`         // 路由配置
}

type RouterConfig struct {
	HealthCheck bool `yaml:"health_check" json:"health_check"`
}

type MiddlewareConfig struct {
	Cors      bool `yaml:"cors" json:"cors"`
	RateLimit bool `yaml:"rateLimit" json:"rateLimit"`
}

type EccConfig struct {
	PrivateKey []byte `yaml:"private_key" json:"private_key"`
	PublicKey  []byte `yaml:"public_key" json:"public_key"`
}
