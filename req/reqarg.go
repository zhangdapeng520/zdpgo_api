package req

import (
	"encoding/json"
	"io"
	"net/http"
	"zdpgo_api/api"
)

// GetPath 获取路径参数
// 定义路径参数语法：/users/{key}
func GetPath(r *http.Request, key string) string {
	return api.URLParam(r, key)
}

func GetPaths(r *http.Request, keys ...string) (values []string) {
	for _, key := range keys {
		values = append(values, GetPath(r, key))
	}
	return
}

// GetHeader 获取请求头参数
func GetHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

func GetHeaders(r *http.Request, keys ...string) (values []string) {
	for _, key := range keys {
		values = append(values, GetHeader(r, key))
	}
	return
}

// GetQuery 获取查询参数
// 定义查询参数的语法：/users?key1=value1&key2=value2
func GetQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetQuerys(r *http.Request, keys ...string) (values []string) {
	for _, key := range keys {
		values = append(values, GetQuery(r, key))
	}
	return
}

// GetForm 获取表单参数
func GetForm(r *http.Request, keys ...string) (values []string) {
	r.ParseForm()
	for _, key := range keys {
		values = append(values, r.PostFormValue(key))
	}
	return
}

// GetJson 获取JSON参数
// 注意：jsonObj要传其指针，比如 GetJson(r, &user)
func GetJson(r *http.Request, jsonObj interface{}) (err error) {
	var body []byte
	body, err = io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, jsonObj)
	return
}
