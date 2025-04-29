package server

import (
	"context"
	"fmt"

	serv "grpc-interceptor/proto/auth"
	"grpc-interceptor/service"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	// 创建grpc服务器，绑定监听端口, 并添加一元中间件
	// srv := grpc.NewServer(
	// 	grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 		start := time.Now()
	// 		// 前置处理：耗时、验证、日志
	// 		resp, err = handler(ctx, req)
	// 		fmt.Println("耗时：", time.Since(start))
	// 		return resp, err
	// 	}),
	// )

	// 创建grpc服务器，绑定监听端口, 并添加链式调用一元中间件
	// srv := grpc.NewServer(
	// 	grpc.ChainUnaryInterceptor(
	// 		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 			start := time.Now()
	// 			// 前置处理：耗时、验证、日志
	// 			resp, err = handler(ctx, req)
	// 			fmt.Println("耗时：", time.Since(start))
	// 			return resp, err
	// 		},
	// 		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 			start := time.Now()
	// 			// 前置处理：耗时、验证、日志
	// 			resp, err = handler(ctx, req)
	// 			fmt.Println("耗时：", time.Since(start))
	// 			return resp, err
	// 		},
	// 	),
	// )

	// 创建grpc服务器，添加官方社区实现的auth拦截器
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
				// 1. 取出token信息
				md, ok := metadata.FromIncomingContext(ctx)
				
				if !ok {
					return ctx, status.Error(codes.Unauthenticated, "noAuthorized")
				}

				token := md.Get("bearer")

				// 2. 验证token
				if tk := token[0]; len(tk) > 0 {
					if tk == "123" {
						return ctx, nil 
					}

				}
				fmt.Println("token")

				return ctx, status.Error(codes.Unauthenticated, "noAuthorized")
			}),
		),
	)
	
	
	// 注册服务
	authServer := &service.AuthServer{}
	serv.RegisterAuthServiceServer(srv, authServer)
	err = srv.Serve(conn)
	if err != nil {
		return err
	}


	return nil
}