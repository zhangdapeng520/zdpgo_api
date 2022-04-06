package main

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 创建核心对象
	g := zdpgo_gin.New(zdpgo_gin.GinConfig{
		Debug: true,
	})

	// 通过map获取json数据
	g.App.POST("/", func(c *gin.Context) {
		json := make(map[string]interface{}) // 注意该结构接受的内容
		c.BindJSON(&json)
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"name": json["name"],
			"age":  json["age"],
		})
	})

	// 通过struct获取json数据
	g.App.PUT("/", func(c *gin.Context) {
		json := Person{}
		c.BindJSON(&json)
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"name":     json.Name,
			"password": json.Age,
		})
	})

	// { "name": "张大鹏" , "age": 22}
	g.Run()
}
