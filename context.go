package zdpgo_api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"io/ioutil"
)

/*
@Time : 2022/5/17 15:50
@Author : 张大鹏
@File : context.go
@Software: Goland2021.3.1
@Description: context上下文对象相关
*/

// Context 重命名gin上下文
type Context struct {
	gin.Context
}

// GetBody 获取body中的字符串数据
func (c *Context) GetBody() ([]byte, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Log.Error("读取Body数据失败", "error", err)
	}
	return body, nil
}

// GetEccBodyToJson 读取ECC加密的内容并解析为JSON数据
func (c *Context) GetEccBodyToJson(jsonData interface{}) error {
	body, err := c.GetBody()
	if err != nil {
		Log.Error("读取请求体内容失败", "error", err)
		return err
	}

	// base64解码
	decodeData, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		Log.Error("解析base64字符串失败", "error", err)
		return err
	}

	// 获取私钥
	ecc := Password.GetEcc()
	privateKey, _, err := ecc.GetKey()
	if err != nil {
		Log.Error("获取私钥失败", "error", err)
		return err
	}

	// ECC解密
	resultData, err := ecc.DecryptByPrivateKey(decodeData, privateKey)
	if err != nil {
		Log.Error("ECC解密数据失败", "error", err)
		return err
	}

	// 解析json数据
	err = json.Unmarshal(resultData, &jsonData)
	if err != nil {
		Log.Error("解析JSON数据失败", "error", err)
		return err
	}

	// 查看结果
	Log.Debug("解析ECC数据成功", "data", jsonData)

	// 返回
	return nil
}
