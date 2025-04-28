package client

import (
	"context"
	"fmt"
	"grpc-interceptor/proto/auth"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() error {
	// 创建客户端, 第二个参数表示不检查证书， 默认会检查证书
	// conn, err := grpc.NewClient("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// if err != nil {
	// 	return err
	// }


	// defer conn.Close()

	// // 绑定
	// client := auth.NewAuthServiceClient(conn)

	// 使用连接池
	cp, err := NewUserClientPool("127.0.0.1:8000", 5)
	if err != nil {
		return nil 
	}
	client := cp.Get()

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 远程调用
	r, err := client.Login(ctx, &auth.LoginRequest{})
	if err != nil {
		return nil 
	}

	fmt.Printf("resp: %s", r.GetUser())
	return nil
}


// 客户端常用池化示例
type userClientPool struct {
	clients []auth.AuthServiceClient
	mu sync.Mutex
	index int64
}

func NewUserClientPool(addr string, size int) (*userClientPool, error) {
	clients := make([]auth.AuthServiceClient, 0, size)
	for range size {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		client := auth.NewAuthServiceClient(conn)

		clients = append(clients, client)
	}

	return &userClientPool{clients: clients, index: 0}, nil
}


func (u *userClientPool) Get() auth.AuthServiceClient {
	// 1. 新增index
	atomic.AddInt64(&u.index, 1)
	return u.clients[int(u.index) % len(u.clients)] // 随着Index增加，只拿到范围内的链接
}

