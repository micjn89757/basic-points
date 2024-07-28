package server

import (
	"fmt"
	"inet/rpc/grpc/stream-grpc/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

// 服务端流模式
func (s *Server) GetStream (request *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {	// 连续向客户端发消息
		i++
		err := res.Send(&proto.StreamResData{	// socket 的send
			Data: fmt.Sprintf("%v", time.Now()),
		}) 
		if err != nil {
			return err
		}

		time.Sleep(time.Second)

		if i > 10 {
			break
		}
	}

	return nil  // 这里虽然会返回，但是server不会关闭
}

// 客户端流模式
func (s *Server) PutStream (clientStream proto.Greeter_PutStreamServer) error {
	for {
		resData, err := clientStream.Recv() // 接收client端的值
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(resData.Data)
	}
	return nil
}


func (s *Server) AllStream (allStream proto.Greeter_AllStreamServer) error {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, err := allStream.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
			fmt.Println("resv client msg:" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := allStream.Send(&proto.StreamResData{Data: "server say hello"})
			time.Sleep(time.Second)
			if err != nil {
				break
			}
		}
	}()
	wg.Wait()
	return nil
}


func (s *Server) Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	serv := grpc.NewServer()
	proto.RegisterGreeterServer(serv, &Server{})
	err = serv.Serve(listener)
	if err != nil {
		panic(err)
	}
}