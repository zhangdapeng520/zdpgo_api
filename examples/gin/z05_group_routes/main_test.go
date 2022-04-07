package main

import (
	"github.com/zhangdapeng520/zdpgo_api/examples/gin/z05_group_routes/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV1PingRoute(t *testing.T) {
	// 创建路由
	router := routes.GetRoutes()

	// 创建记录器
	w := httptest.NewRecorder()

	// 创建请求
	req, _ := http.NewRequest("GET", "/v1/ping/", nil)

	// 发送请求
	router.ServeHTTP(w, req)

	// 断言ping路径
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestV1UsersRoute(t *testing.T) {
	// 创建路由
	router := routes.GetRoutes()

	// 创建记录器
	w := httptest.NewRecorder()

	// 断言 /users/路径
	req, _ := http.NewRequest("GET", "/v1/users/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "users", w.Body.String())
}

func TestV1UsersCommentsRoute(t *testing.T) {
	// 创建路由
	router := routes.GetRoutes()

	// 创建记录器
	w := httptest.NewRecorder()

	// 断言 /users/comments路径
	req, _ := http.NewRequest("GET", "/v1/users/comments/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "users comments", w.Body.String())
}

func TestV1UsersPicturesRoute(t *testing.T) {
	// 创建路由
	router := routes.GetRoutes()

	// 创建记录器
	w := httptest.NewRecorder()

	// 断言 /users/pictures路径
	req, _ := http.NewRequest("GET", "/v1/users/pictures/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "users pictures", w.Body.String())
}

func TestV2PingRoute(t *testing.T) {
	// 创建路由
	router := routes.GetRoutes()

	// 创建记录器
	w := httptest.NewRecorder()

	// 断言 /users/pictures路径
	req, _ := http.NewRequest("GET", "/v2/ping/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
