package server

import (
	"grpc-starter/proto/auth"
	"grpc-starter/service"
	"net"

	"google.golang.org/grpc"
)

func NewGRPCServer() error {
	// 创建监听端口
    tcpAdd, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		return err 
	}

	conn, err := net.ListenTCP("tcp", tcpAdd)
	if err != nil {
		return err 
	}

	// 创建grpc服务器，绑定监听端口
	srv := grpc.NewServer()
	// 注册服务
	authServer := &service.AuthServer{}
	auth.RegisterAuthServiceServer(srv, authServer)
	err = srv.Serve(conn)
	if err != nil {
		return err
	}


	return nil
}