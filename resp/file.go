package resp

import (
	"net/http"
)

// File 响应文件
func File(w http.ResponseWriter, content []byte) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(content)
}
