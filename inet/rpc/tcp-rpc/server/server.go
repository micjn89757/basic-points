package main

import (
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	Serv *Service 
}

func NewServer(serv *Service) *Server {
	return &Server{
		Serv: serv,
	}
}

func (s *Server) Run() {
	rpc.Register(s.Serv) // 注册为一个rpc服务
	// rpc.HandleHTTP()	// 基于HTTP协议
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("Listen error", err)
	}

	for {	// 基于tcp
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}

func main() {
	Server := NewServer(&Service{})
	Server.Run()
}