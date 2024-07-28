package main

import (
	"context"
	"inet/rpc/grpc/hello-grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	
	defer conn.Close()

	Client := proto.NewGreeterClient(conn)


	// 执行RPC调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := Client.SayHello(ctx, &proto.HelloRequest{Name: "demo", Age: 12, Courses: []string{"1", "2"}, Mp: map[string]string{"1":"2"}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetReply())

}