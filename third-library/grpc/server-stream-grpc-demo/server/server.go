package main

import (
	"context"
	"log"
	"net"
	pb "server-stream-grpc-demo/proto"
	"strconv"

	"google.golang.org/grpc"
)

// StreamService 定义服务
type StreamService struct {
	pb.UnimplementedStreamServerServer
}

const (
	// address 监听地址
	Address string = ":9090"
	// network 网络通信协议
	Network string = "tcp"
)

// Route 实现 Route 方法
func (s *StreamService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

// ListValue 实现 ListValue 方法
func (s *StreamService) ListValue(req *pb.SimpleRequest, serv pb.StreamServer_ListValueServer) error {

	for n := 0; n < 5; n++ {
		// 向流中发送消息，默认每次 send 送消息的最大长度为 `math.MaxInt32`bytes
		err := serv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}
	log.Println(Address + "net.Listen")
	// 新建 gRPC 服务器实例
	grpcServer := grpc.NewServer()
	// 在 gRPC 服务器中注册我们的服务
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})
	// 启动服务，用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
