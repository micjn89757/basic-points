package server

import (
	"fmt"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"grpc-gateway_/internal/config"
	"grpc-gateway_/internal/server/middleware"
	"grpc-gateway_/internal/service"
	pb "grpc-gateway_/proto/mypb"
)

// Run 启动rpc服务
func Run() error {
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	grpcAddr := config.GetRpcAddr()
	// 127.0.0.1:8001
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		grpclog.Errorf("tcp listen failed: %v", err)
		return err
	}
	defer func() {
		if err = l.Close(); err != nil {

			fmt.Fprintln(os.Stderr, err)
		}
	}()

	op := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	}

	s := grpc.NewServer(op...)

	// 	注册服务
	registerServer(s)

	log.Infof("Serving gRPC on %s", l.Addr())

	return s.Serve(l)
}

func registerServer(s *grpc.Server) {
	pb.RegisterBookServiceServer(s, service.NewBookService())

}