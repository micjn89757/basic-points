package main

import (
	"context"
	"fmt"
	"inet/rpc/grpc/grpc-validate/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)


func main() {
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	
	defer conn.Close()

	Client := proto.NewGreeterClient(conn)


	// 超时控制，传入ctx即可
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	r, err := Client.SayHello(ctx, &proto.Person{
		Id: 1000,
		Mobile: "156511651",
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic(ok)
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}

	log.Printf("Greeting: %s", r.GetEmail())

}