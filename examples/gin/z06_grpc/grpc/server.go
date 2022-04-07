package main

import (
	pb "github.com/gin-gonic/examples/grpc/pb"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server 实现grpc服务
type server struct{}

// SayHello 实现发送消息的接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 创建监听
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("创建监听失败: %v", err)
	}

	// 创建grpc服务
	s := grpc.NewServer()

	// 注册消息服务
	pb.RegisterGreeterServer(s, &server{})

	// 注册服务
	reflection.Register(s)

	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动grpc服务失败: %v", err)
	}
}
