package main

import (
	"fmt"
	pb "github.com/zhangdapeng520/zdpgo_api/examples/z06_grpc/pb"
	"github.com/zhangdapeng520/zdpgo_api/gin"
	"log"
	"net/http"

	"google.golang.org/grpc"
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

func getRouter() *gin.Engine {
	// 创建服务
	r := gin.Default()

	// 获取grpc客户端
	client := getGrpcClient()

	// 监听路径
	r.GET("/rest/n/:name", func(c *gin.Context) {

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

	return r
}

func main() {
	// 获取router
	r := getRouter()

	// 启动服务
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
