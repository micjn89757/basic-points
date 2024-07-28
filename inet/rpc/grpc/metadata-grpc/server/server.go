package server

import (
	"context"
	"fmt"
	"inet/rpc/grpc/hello-grpc/proto"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct{
	proto.UnimplementedGreeterServer
}


func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata: ", ok)
	}

	for key, val := range md {
		fmt.Println(key, val)
	}

	// fmt.Println(request.Name)

	return &proto.HelloResponse{
		Reply: "Name:" + request.Name + ", Age:" + strconv.FormatInt(int64(request.Age), 10),
		G: proto.Gender_FEMALE,	// 这样使用枚举类型
	}, nil
}


func (s *Server) Run() {
	grpcServ := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServ, &Server{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = grpcServ.Serve(listener)  // 不会接收到一个消息就退出，实际上每来一个协程就退出

	if err != nil {
		panic("failed to start")
	}
}