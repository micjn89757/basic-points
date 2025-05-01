package service

import (
	pb "grpc-stream/proto"
	"io"

	"google.golang.org/grpc"
)

func NewBlogService() *BlogService {
	return &BlogService{}
}

// 服务端流
func (s *BlogService) LotsOfReplies(req *pb.Request, stream grpc.ServerStreamingServer[pb.Response]) error {
	words := []string{
		"hello",
		"h",
		"D",
	}


	for _, v := range words {
		data := &pb.Response{
			Reply: v + req.Name,
		}
		// 连续返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}

	return nil 
}

// 客户端流
func (s *BlogService) LotsOfReplies1(stream grpc.ClientStreamingServer[pb.Request, pb.Response]) error {
	reply := "hi"

	for {
		// 接受客户端发来的流数据
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Reply: reply,
			})
		}

		if err != nil {
			return err 
		}

		reply += res.GetName()
	}
}

// 双流
func (s *BlogService) LotsOfReplies2(stream grpc.BidiStreamingServer[pb.Request, pb.Response]) error {
	for {
		// 接收客户端流
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err 
		}

		reply := in.GetName() // 处理

		// 返回流式响应
		if err := stream.Send(&pb.Response{Reply: reply}); err != nil {
			return err
		}
	}
}