package main

import (
	"context"
	"fmt"
	"inet/rpc/grpc/hello-grpc/pb"
	"net"

	"google.golang.org/grpc"
)

// 实现rpc服务
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "hello" + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(listen)
	if err != nil {
		return 
	}
}