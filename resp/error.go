package resp

import (
	"encoding/json"
	"net/http"
)

// ErrorMap 响应错误的JSON类型的数据
func ErrorMap(w http.ResponseWriter, statusCode int, kvs ...interface{}) {
	jsonData := make(map[string]interface{})
	if len(kvs) >= 2 {
		for i := 0; i < len(kvs); i += 2 {
			if k, ok := kvs[i].(string); ok {
				jsonData[k] = kvs[i+1]
			}
		}
	} else {
		jsonData["error"] = "服务端错误"
	}
	v, _ := json.Marshal(jsonData)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}

// ErrorMessage 响应错误信息
func ErrorMessage(w http.ResponseWriter, msg string) {
	ErrorMap(w, 200, "status", false, "code", 1001, "msg", msg)
}

// ErrorCodeMessage 响应错误的状态码和错误信息
func ErrorCodeMessage(w http.ResponseWriter, code int, msg string) {
	ErrorMap(w, 200, "status", false, "code", code, "msg", msg)
}
