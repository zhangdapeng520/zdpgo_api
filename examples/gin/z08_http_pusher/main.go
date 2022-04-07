package main

import (
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
	"html/template"
	"log"
	"net/http"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	// 创建路由
	r := gin.Default()

	// 设置静态文件路径
	r.Static("/assets", "examples/gin/z08_http_pusher/assets")

	// 设置模板
	r.SetHTMLTemplate(html)

	// 监听路径
	r.GET("/", func(c *gin.Context) {
		// 创建pusher
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用pusher.Push()实现服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("推送失败: %v", err)
			}
		}
		// 返回HTML响应
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// 监听路径，会解析模板，拉取js代码
	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并启动服务 https://127.0.0.1:8080
	// 监听并启动服务 https://127.0.0.1:8080/welcome
	r.RunTLS(":8080", "examples/gin/z08_http_pusher/testdata/server.pem", "examples/gin/z08_http_pusher/testdata/server.key")
}
