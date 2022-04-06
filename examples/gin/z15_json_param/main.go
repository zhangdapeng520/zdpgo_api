package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	route := gin.Default()

	// 通过map获取json数据
	route.POST("/", func(c *gin.Context) {
		json := make(map[string]interface{}) // 注意该结构接受的内容
		c.BindJSON(&json)
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"name": json["name"],
			"age":  json["age"],
		})
	})

	// 通过struct获取json数据
	route.PUT("/", func(c *gin.Context) {
		json := Person{}
		c.BindJSON(&json)
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"name":     json.Name,
			"password": json.Age,
		})
	})

	// { "name": "张大鹏" , "age": 22}
	route.Run(":8080")
}
