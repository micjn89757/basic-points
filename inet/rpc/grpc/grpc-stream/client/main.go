package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// !这里不能用localhost:8000
	conn, err := grpc.NewClient("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("did not connect")
	}


	defer conn.Close()


	client := pb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.LotsOfReplies(ctx, &pb.Request{Name: "d"})

	if err != nil {
		log.Fatal("did not connect")
	}

	for {
		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
		res, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("failed")
		}

		log.Println(res.GetReply())
	}
}