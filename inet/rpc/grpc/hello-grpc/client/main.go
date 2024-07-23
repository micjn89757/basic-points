package main

import (
	"context"
	"flag"
	"inet/rpc/grpc/hello-grpc/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to ")
	name = flag.String("name", defaultName, "name to great")
)

func main() {
	flag.Parse()

	// 连接到server端，此处禁用安全传输
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()

	Client := pb.NewGreeterClient(conn)


	// 执行RPC调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := Client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetReply())

}