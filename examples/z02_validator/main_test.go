package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookableRoute(t *testing.T) {
	// 创建路由
	router := setupRouter()

	// 创建记录器
	w := httptest.NewRecorder()

	// 创建请求
	req, _ := http.NewRequest("GET", "/bookable?check_in=2022-05-02&check_out=2022-05-03", nil)

	// 发送请求
	router.ServeHTTP(w, req)

	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 查看响应
	var data = make(map[string]string)
	t.Log(w.Body.String())

	// 断言响应内容
	_ = json.Unmarshal([]byte(w.Body.String()), &data)
	assert.Equal(t, "预订日期有效！", data["message"])
}
