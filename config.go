package zdpgo_api

/*
@Time : 2022/5/16 16:25
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: config配置相关
*/

type Config struct {
	Debug       bool   `yaml:"debug" json:"debug"`                 // 是否为debug模式
	LogFilePath string `yaml:"log_file_path" json:"log_file_path"` // 日志路径
	Host        string `yaml:"host" json:"host"`                   // 启动地址，默认0.0.0.0
	Port        int    `yaml:"port" json:"port"`                   // 启动端口号，默认3333
}
