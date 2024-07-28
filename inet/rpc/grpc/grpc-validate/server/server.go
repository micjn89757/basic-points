package server

import (
	"context"
	"inet/rpc/grpc/grpc-validate/proto"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{
	proto.UnimplementedGreeterServer
}


func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person, error) {
	time.Sleep(time.Second * 5)
	return &proto.Person{Id: 5000}, status.Errorf(codes.NotFound, "invalid email:%s", request.Email)
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