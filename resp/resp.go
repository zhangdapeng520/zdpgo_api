package resp

import (
	"encoding/json"
	"net/http"
)

// Json 响应JSON数据
func Json(w http.ResponseWriter, jsonData interface{}) {
	v, _ := json.Marshal(jsonData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}

// JsonMap 响应JSON类型的数据
// 注意：kvs必须是键值对的形式，且key必须是字符串
func JsonMap(w http.ResponseWriter, kvs ...interface{}) {
	jsonData := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		if k, ok := kvs[i].(string); ok {
			jsonData[k] = kvs[i+1]
		}
	}
	v, _ := json.Marshal(jsonData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}
