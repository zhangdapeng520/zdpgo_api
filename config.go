package zdpgo_api

/*
@Time : 2022/5/16 16:25
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: config配置相关
*/

type Config struct {
	Debug          bool      `yaml:"debug" json:"debug"`
	Host           string    `yaml:"host" json:"host"`                         // 启动地址，默认0.0.0.0
	Port           int       `yaml:"port" json:"port"`                         // 启动端口号，默认3333
	UploadFileSize int64     `yaml:"upload_file_size" json:"upload_file_size"` // 上传文件大小限制（M），默认33
	Ecc            EccConfig `yaml:"ecc" json:"ecc"`
}

type EccConfig struct {
	PrivateKey []byte `yaml:"private_key" json:"private_key"`
	PublicKey  []byte `yaml:"public_key" json:"public_key"`
}
