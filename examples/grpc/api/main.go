package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_api"
	pb "github.com/zhangdapeng520/zdpgo_api/examples/grpc/pb"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func getGrpcClient() pb.GreeterClient {
	// 创建连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("建立GRPC连接失败: %v", err)
	}

	// 使用优雅退出，在优雅退出时释放此资源
	// defer conn.Close()

	// 创建客户端
	client := pb.NewGreeterClient(conn)
	return client
}

func main() {
	// 创建服务
	api := zdpgo_api.NewWithConfig(zdpgo_api.Config{Debug: true})

	// 获取grpc客户端
	client := getGrpcClient()

	// 监听路径
	api.App.GET("/rest/n/:name", func(c *gin.Context) {

		// 获取名字
		name := c.Param("name")

		// 创建请求
		req := &pb.HelloRequest{Name: name}

		// 发送请求
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 返回响应
		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	// 启动服务
	// 测试：http://127.0.0.1:3333/rest/n/zhangdapeng
	api.Run()
}
