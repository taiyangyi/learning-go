package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "simple-grpc-demo/proto/helloworld"
	"simple-grpc-demo/service/helloworld"
)

const (
	host      = "127.0.0.1"
	grpc_port = 9090
	network   = "tcp"
)

func main() {
	address := fmt.Sprintf("%s:%d", host, grpc_port)
	listen, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 初始化 gRPC 服务
	server := grpc.NewServer()
	// 将 gRPC 服务和自定义的业务逻辑注册到 Greeter 服务中
	pb.RegisterGreeterServer(server, &helloworld.Server{})
	log.Printf("server gRPC on %v", listen.Addr())
	// 将 gRPC 服务绑定在上面创建的 TCP 端口上并开启监听
	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
