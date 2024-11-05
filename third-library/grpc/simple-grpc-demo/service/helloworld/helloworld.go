package helloworld

import (
	"context"
	"log"
	pb "simple-grpc-demo/proto/helloworld"
)

// Server 创建 Server 结构体
type Server struct {
	// Server 结构体嵌入 pb.UnimplementedGreeterServer 结构体的所有方法
	pb.UnimplementedGreeterServer
}

// SayHello 重写 pb.GreeterServer.SayHello 方法，实现业务逻辑
func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received message from %v", request.GetName())
	return &pb.HelloResponse{Message: "Hello " + request.GetName()}, nil
}
