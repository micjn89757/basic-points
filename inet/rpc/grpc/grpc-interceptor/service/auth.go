package service

import (
	"grpc-interceptor/proto/auth"
	"context"
)

type AuthServer struct {
	// 嵌入的server提供默认的方法实现
	auth.UnimplementedAuthServiceServer  // 每一个Grpc服务都必须嵌入， 嵌入时使用值而不是指针类型结构体，防止空指针
}

// 参数验证
func validateLoginRequest(request *auth.LoginRequest) error {
	return nil
}

func validateRegisterRequest(request *auth.RegisterRequest) error {
	return nil 
}

func (as *AuthServer) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	if err := validateLoginRequest(request); err != nil {
		return nil, err 
	}

	// biz

	// 做完了业务，组装响应数据

	// 返回参数

	return &auth.LoginResponse{
		User: &auth.User{
			Name: "ddd",
			Password: "asdjkl",
		},
		Token: "addddddd",
	}, nil
}

func (as *AuthServer) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if err := validateRegisterRequest(request); err != nil {
		return nil, err 
	}


	return nil, nil
}