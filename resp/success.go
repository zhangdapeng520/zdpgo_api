package resp

import (
	"encoding/json"
	"net/http"
)

// SuccessMap 响应成功的数据，以map的形式指定data
func SuccessMap(w http.ResponseWriter, kvs ...interface{}) {
	// 基本数据结构
	jsonData := map[string]interface{}{
		"status": true,
		"code":   10000,
		"msg":    "success",
	}

	// 附加的数据
	if len(kvs) >= 2 {
		data := make(map[string]interface{})
		for i := 0; i < len(kvs); i += 2 {
			if k, ok := kvs[i].(string); ok {
				data[k] = kvs[i+1]
			}
		}
		jsonData["data"] = data
	}

	// 返回
	v, _ := json.Marshal(jsonData)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}

// Success 响应成功的数据，data是一个JSON对象
func Success(w http.ResponseWriter, data interface{}) {
	// 基本数据结构
	jsonData := map[string]interface{}{
		"status": true,
		"code":   10000,
		"msg":    "success",
	}

	// 附加的数据
	if data != nil {
		jsonData["data"] = data
	}

	// 返回
	v, _ := json.Marshal(jsonData)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}
