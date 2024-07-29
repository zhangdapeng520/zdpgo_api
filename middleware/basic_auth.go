package middleware

import (
	"crypto/subtle"
	"fmt"
	"net/http"
)

// BasicAuth 实现了一个简单的中间件处理程序，用于向路由添加基本的 http 验证。
func BasicAuth(realm string, creds map[string]string) func(next http.Handler) http.Handler {
	// 返回一个中间件
	return func(next http.Handler) http.Handler {
		// 返回一个处理器
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 获取用户名和密码
			user, pass, ok := r.BasicAuth()
			if !ok {
				basicAuthFailed(w, realm)
				return
			}

			// 根据用户名获取密码并进行校验
			credPass, credUserOk := creds[user]
			if !credUserOk || subtle.ConstantTimeCompare([]byte(pass), []byte(credPass)) != 1 {
				basicAuthFailed(w, realm)
				return
			}

			// 校验通过，正常执行后面的逻辑
			next.ServeHTTP(w, r)
		})
	}
}

// 基本校验失败
func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
