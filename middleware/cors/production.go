package cors

import (
	"net/http"
)

var _origins = new([]string)

// Production 返回生产环境的跨域中间件，允许任何域名
func Production(origin string, origins ...string) func(next http.Handler) http.Handler {
	origins = append(origins, origin)
	*_origins = origins
	return Handler(Options{
		AllowOriginFunc:  allowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

func allowOriginFunc(r *http.Request, origin string) bool {
	for _, _origin := range *_origins {
		if _origin == origin {
			return true
		}
	}
	//if origin == "http://localhost:8888" {
	//	return true
	//}
	return false
}
