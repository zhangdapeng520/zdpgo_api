package main

import (
	"github.com/zhangdapeng520/zdpgo_api/libs/gin"
	"html/template"
	"log"
	"net/http"
	"os"
)

// 使用https解析模板
var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	// 创建日志
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	// 创建路由
	r := gin.Default()

	// 设置模板
	r.SetHTMLTemplate(html)

	// 监听路径
	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// 启动服务 https://127.0.0.1:8080
	// 启动服务 https://127.0.0.1:8080/welcome
	r.RunTLS(":8080", "examples/gin/z07_http2/testdata/server.pem", "examples/gin/z07_http2/testdata/server.key")
}
