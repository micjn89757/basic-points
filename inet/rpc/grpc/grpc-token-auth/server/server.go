package server

import (
	"context"
	"fmt"
	"inet/rpc/grpc/grpc-token-auth/proto"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct{
	proto.UnimplementedGreeterServer
}


func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
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
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			// 已经开始接触到grpc的错误处理了
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}


		var (
			appid string 
			appkey string 
		)

		if val, ok := md["appid"]; ok {
			appid=val[0]
		}

		if val, ok := md["appkey"]; ok {
			appkey=val[0]
		}

		if appid != "bobby" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
	
	
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