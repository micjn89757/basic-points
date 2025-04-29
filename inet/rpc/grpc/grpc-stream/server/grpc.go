package server

import (

	pb "grpc-stream/proto"
	"grpc-stream/service"

	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {

	srv := grpc.NewServer() // 创建grpc服务器

	pb.RegisterBlogServiceServer(srv, &service.BlogService{}) // 注册服务

	return srv
}