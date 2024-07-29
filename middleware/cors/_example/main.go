package main

import (
	"net/http"
	"zdpgo_api/api"
	"zdpgo_api/middleware"
	"zdpgo_api/middleware/cors"
)

func main() {
	r := api.NewRouter()
	r.Use(middleware.Logger)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":3000", r)
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://example.com" {
		return false
	}
	return true
}
