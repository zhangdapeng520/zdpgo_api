package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	// 创建路由
	router := setupRouter()

	// 创建记录器
	w := httptest.NewRecorder()

	// 创建请求
	req, _ := http.NewRequest("GET", "/ping", nil)

	// 发送请求
	router.ServeHTTP(w, req)

	// 断言
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// 测试用户路由
func TestUserRoute(t *testing.T) {
	// 创建路由
	router := setupRouter()

	// 创建记录器
	w := httptest.NewRecorder()

	// 创建请求
	req, _ := http.NewRequest("GET", "/user/zhangdapeng520", nil)

	// 发送请求
	router.ServeHTTP(w, req)

	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 获取返回值并断言
	var data = make(map[string]string)
	_ = json.Unmarshal([]byte(w.Body.String()), &data)

	t.Log("返回值是：", data)
	t.Log("返回值是：", w.Body.String())

	assert.Equal(t, data["user"], "zhangdapeng520")
	assert.Equal(t, data["status"], "no value")
}
