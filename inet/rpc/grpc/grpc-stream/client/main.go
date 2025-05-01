package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	// ServerStream(client) // 服务端流
	// ClientStream(client) // 客户端流
	ClientServerStream(client) // 双向流

}


func ServerStream(client pb.BlogServiceClient) {
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

func ClientStream(client pb.BlogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := client.LotsOfReplies1(ctx)
	if err != nil {
		log.Fatal("err")
	}


	names := []string{"qimi", "dd", "ff"}

	for _, name := range names {
		err := stream.Send(&pb.Request{Name: name}) // 发送数据
		if err != nil {
			log.Fatal("err")
		}
	}
	res, _ := stream.CloseAndRecv()
	log.Println("get reply:", res.GetReply())
}


func ClientServerStream(client pb.BlogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Minute)
	defer cancel()

	// 双向流
	stream, err := client.LotsOfReplies2(ctx)

	if err != nil {
		log.Fatal(err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			// 接收服务端返回的数据 并打印
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(in.Reply)
		}
	}()

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin)

	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)

		if len(cmd) == 0 {
			continue
		}

		if strings.ToUpper(cmd) == "Q" {
			break
		}

		// 将获取到的数据发送到服务端
		if err := stream.Send(&pb.Request{Name: cmd}); err != nil {
			log.Fatal(err)
		}
	}

	stream.CloseSend()
	<-waitc

}