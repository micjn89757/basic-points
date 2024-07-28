package server

import (
	"context"
	"fmt"
	"inet/rpc/grpc/grpc-interpretor/proto"
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
	// grpc普通模式的server拦截器实现，当然还有流模式的实现
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) { 
		fmt.Println("接收到了一个新请求")
		return handler(ctx, req)  // 继续后面的逻辑
	}
	opt := grpc.UnaryInterceptor(interceptor)
	grpcServ := grpc.NewServer(opt)
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