package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestNRoute(t *testing.T) {
	// 创建路由
	router := getRouter()

	// 创建记录器
	w := httptest.NewRecorder()

	// 创建请求
	req, _ := http.NewRequest("GET", "/rest/n/zhangdapeng520", nil)

	// 发送请求
	router.ServeHTTP(w, req)
	t.Log("响应：", w.Body.String())

	// 断言ping路径
	assert.Equal(t, http.StatusOK, w.Code)
	var data = make(map[string]string)
	_ = json.Unmarshal([]byte(w.Body.String()), &data)
	assert.Equal(t, "Hello zhangdapeng520", data["result"])
}
