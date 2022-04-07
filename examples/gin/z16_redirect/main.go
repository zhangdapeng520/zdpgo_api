package main

import (
	"net/http"

	"zdpgo_api"
)

func main() {
	route := zdpgo_api.Default()

	route.GET("/test", func(c *zdpgo_api.Context) {
		// c.Request.URL.Path = "/test2"
		// route.HandleContext(c)
		c.Redirect(http.StatusMovedPermanently, "/test2")
	})

	route.GET("/test2", func(c *zdpgo_api.Context) {
		c.JSON(200, zdpgo_api.H{"hello": "world"})
	})

	// { "name": "张大鹏" , "age": 22}
	route.Run(":8080")
}
