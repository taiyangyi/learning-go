package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/hello-server/proto"
	"log"
)

func main() {

	// 连接到 server 端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行 rpc 调用（这个方法在服务器来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "bamaw"})
	fmt.Println(resp.GetResponseMsg())
}
