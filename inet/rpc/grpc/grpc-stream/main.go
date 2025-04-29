package main

import (
	"grpc-stream/server"
	"net"
)

func main() {
	// 创建端口
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8000")

	conn, _ := net.ListenTCP("tcp", tcpAddr)

	srv := server.NewGRPCServer()

	srv.Serve(conn)
}