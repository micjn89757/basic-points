package main

import (
	"context"
	"inet/rpc/grpc/hello-grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)


func main() {
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	
	defer conn.Close()

	Client := proto.NewGreeterClient(conn)

	// 添加metadata
	// 使用Pars方式
	// md := metadata.Pairs("password", "immoc")

	// 使用new的方式
	md := metadata.New(map[string]string{
		"name": "bobby",
		"password": "immoc",
	})

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// 将metadata添加到上下文
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// fmt.Printf("%#v", ctx)

	// 执行RPC调用
	r, err := Client.SayHello(ctx, &proto.HelloRequest{Name: "demo", Age: 12, Courses: []string{"1", "2"}, Mp: map[string]string{"1":"2"}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetReply())

}