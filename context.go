package zdpgo_api

import (
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

func (c *Context) GetAesTextBodyToJson(jsonData interface{}) error {
	body, err := c.GetBody()
	if err != nil {
		Log.Error("读取请求体内容失败", "error", err)
		return err
	}

	// AES解密
	resultData, err := Password.Aes.Decrypt(body)
	if err != nil {
		Log.Error("AES解密数据失败", "error", err, "body", body)
		return err
	}

	// 解析json数据
	err = json.Unmarshal(resultData, &jsonData)
	if err != nil {
		Log.Error("解析JSON数据失败", "error", err, "resultData", resultData)
		return err
	}

	// 返回
	return nil
}

func (c *Context) ResponseAesStr(jsonResponse interface{}) {
	var result string

	// 将结果转换为JSON字符串
	jsonStrBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		result = err.Error()
		c.String(501, result)
		return
	}

	// 加密结果数据
	aesBytes, err := Password.Aes.Encrypt(jsonStrBytes)
	if err != nil {
		result = err.Error()
		c.String(501, result)
		return
	}

	// 返回加密数据
	c.String(200, string(aesBytes))
}
