package service

import  pb "grpc-stream/proto"

// 定义服务
type BlogService struct {
	pb.UnimplementedBlogServiceServer
}