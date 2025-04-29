package service 


import (
	pb "grpc-stream/proto"
	"google.golang.org/grpc"
)

func NewBlogService() *BlogService {
	return &BlogService{}
}


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