package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	X, Y int
}

func main() {
	// 建立tcp连接
	client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	// client, err := rpc.DialHttp("tcp", "127.0.0.1:9091") // 建立http连接
	if err != nil {
		log.Fatal("dialing:", err)
	}


	// 同步调用
	args := &Args{
		10, 
		20,
	}

	var reply int
	err = client.Call("Service.Add", args, &reply)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Add: %d+%d=%d\n", args.X, args.Y, reply)


	// 异步调用
	var reply2 int 
	divCall	:= client.Go("Service.Add", args, &reply2, nil)

	replyCall := <- divCall.Done	// 接收调用结果
	fmt.Printf("%#v", replyCall)
	fmt.Println(reply2)
}