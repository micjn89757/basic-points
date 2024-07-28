package server

import (
	"context"
	"inet/rpc/grpc/grpc-validate/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{
	proto.UnimplementedGreeterServer
}


func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person, error) {
	return &proto.Person{Id: 5000}, status.Errorf(codes.NotFound, "invalid email:%s", request.Email)
}

type Validator interface { // protoc-gen-validate会为每一个使用校验的消息创建Validate方法
	Validate() error
}

func (s *Server) Run() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
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