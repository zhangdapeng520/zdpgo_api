package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
)

func main() {
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{Debug: true})
	api.AddUploadRouter("/upload", "file", "uploads")
	api.Run()
}
