package main

import (
	"github.com/zhangdapeng520/zdpgo_api"
	"github.com/zhangdapeng520/zdpgo_log"
)

/*
@Time : 2022/6/6 15:14
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	privateKey := `-----BEGIN  ZDPGO_PASSWORD ECC PRIVATE KEY -----
MHcCAQEEIKyfOnD7NdXudekftRtH2mBuOPf/UTzJ1Ulo2Hiu22XvoAoGCCqGSM49
AwEHoUQDQgAEXClGdjDvOFSHJzs2LtSfGcVzP58cc9ybrYOo7t6bs818HMybbahM
Qylb+qB4aTtHV0JPqZAr8MChRmvze7nNFw==
-----END  ZDPGO_PASSWORD ECC PRIVATE KEY -----
`
	publicKey := `-----BEGIN  ZDPGO_PASSWORD ECC PUBLIC KEY -----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXClGdjDvOFSHJzs2LtSfGcVzP58c
c9ybrYOo7t6bs818HMybbahMQylb+qB4aTtHV0JPqZAr8MChRmvze7nNFw==
-----END  ZDPGO_PASSWORD ECC PUBLIC KEY -----`

	api := zdpgo_api.NewWithConfig(&zdpgo_api.Config{
		Ecc: zdpgo_api.EccConfig{
			PrivateKey: []byte(privateKey),
			PublicKey:  []byte(publicKey),
		},
	}, zdpgo_log.NewWithDebug(true, "log.log"))

	api.Post("/ecctext", func(ctx *zdpgo_api.Context) {
		// 解析json数据
		var jsonData struct {
			Username string `json:"username"`
			Age      int    `json:"age"`
		}
		err := ctx.GetEccTextBodyToJson(&jsonData)
		if err != nil {
			panic(err)
		}

		// 加密响应数据
		response := &zdpgo_api.Response{
			Code:   10000,
			Msg:    "success",
			Status: true,
			Data:   jsonData,
		}
		ctx.ResponseEccStr(api, response)
	})

	// 启动
	api.Run()
}
