package main

import (
	"context"
	"fmt"
	"inet/rpc/grpc/stream-grpc/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := proto.NewGreeterClient(conn)

	// // GetStream
	// res, _ := client.GetStream(context.Background(), &proto.StreamReqData{
	// 	Data: "client request",
	// })

	// for {
	// 	resData, err := res.Recv()	// socket编程中的send recv
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}

	// 	fmt.Printf("client resv Data:%v\n",resData.Data)
	// }

	// PutStream
	// putS, _ := client.PutStream(context.Background())
	// i := 0 
	// for {
	// 	i++
	// 	err := putS.Send(&proto.StreamReqData{Data: "hello"})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	time.Sleep(time.Second)

	// 	if i > 10 {
	// 		break
	// 	}
	// }


	allS, _ := client.AllStream(context.Background())
	i := 0 
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			i++
			err := allS.Send(&proto.StreamReqData{Data: "hello"})
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)

			if i > 10 {
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			data, err := allS.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
			fmt.Printf("client resv data:%s\n", data.Data)
		}
	}()
	wg.Wait()
}