package middleware

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var xUserId = "x-user-id"

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// md, _ := metadata.FromIncomingContext(ctx)
	
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md["x-user-id"]) == 0 {
		return nil, errors.New("UnAuthorization")
	}
	
	userId := strings.Join(md[xUserId], ",")
	ctx = context.WithValue(ctx, xUserId, userId)
	return handler(ctx, req)
}
