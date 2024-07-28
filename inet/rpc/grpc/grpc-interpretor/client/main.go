package main

import (
	"context"
	"fmt"
	"inet/rpc/grpc/grpc-interpretor/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)


func main() {
	// 客户端拦截器
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error { // 客户端的后期逻辑放在invoker中
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时: %s\n", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	// conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)

	// NewClient第二个参数的要求可以改成如下
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, opt)
	conn, err := grpc.NewClient("127.0.0.1:8080", opts...)

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