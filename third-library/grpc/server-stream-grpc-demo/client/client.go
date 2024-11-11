package main

import (
	"context"
	"io"
	"log"
	pb "server-stream-grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Address 连接地址
const Address string = ":9090"

var grpcClient pb.StreamServerClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立 gRPC 连接
	grpcClient = pb.NewStreamServerClient(conn)
	route()
	listValue()
}

// route 调用服务端 Route 方法
func route() {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc",
	}

	// 调用我们的服务（Route）方法
	// 同时传入一个 context.Context，在有需要时可以让我们改变 RPC 的行为，比如：超时/取消一个正在运行的 RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("call Route err: %v", err)
	}
	log.Println(res)
}

// listValue 调用服务端的 ListValue 方法
func listValue() {
	// 创建发送的结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc",
	}

	// 调用我们的服务（ListValue）方法
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("call listStr err: %v", err)
	}
	for {

		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息刘是否已经结束
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}

		// 打印返回值
		log.Println(res.StreamValue)

	}
	// 可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	// 调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	// stream.CloseSend()
}
