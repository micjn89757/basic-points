package gateway

import (
	"context"
	// "io"
	"net/http"
	// "os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/grpc/grpclog"

	"grpc-gateway_/internal/config"
	"grpc-gateway_/internal/gateway/middleware"
	"grpc-gateway_/internal/handler"
	pb "grpc-gateway_/proto/mypb"
)

func Run() error {
	// log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	// grpclog.SetLoggerV2(log)

	ctx := context.Background()

	op := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(middleware.AuthInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// 创建grpc连接 127.0.0.1:8001
	// conn, err := grpc.DialContext(ctx, config.GetRpcAddr(), op...)
	conn, err := grpc.NewClient(config.GetRpcAddr(), op...)
	if err != nil {
		// log.Error("dial failed: %v", err)
		return err
	}

	sp := []runtime.ServeMuxOption{
		runtime.WithForwardResponseOption(middleware.Forward),
		runtime.WithRoutingErrorHandler(middleware.RoutingErrorHandler),
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			if s == "X-User-Id" {
				return s, true
			}
			return runtime.DefaultHeaderMatcher(s)
		}),
	}
	mux := runtime.NewServeMux(sp...)

	// 文件上传不适合使用proto定义，可以自定义路由
	if err = mux.HandlePath(http.MethodPost, "/v1/objects", handler.Upload); err != nil {
		return err
	}

	if err = mux.HandlePath(http.MethodGet, "/v1/objects/{name}", handler.Download); err != nil {
		return err
	}

	err = newGateway(ctx, conn, mux)
	if err != nil {
		// log.Error("register handler failed: %v", err)
		return err
	}

	server := http.Server{
		Addr:    config.GetHttpAddr(), // 127.0.0.1:8002
		Handler: mux,
	}

	// log.Infof("Serving Http on %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func newGateway(ctx context.Context, conn *grpc.ClientConn, mux *runtime.ServeMux) error {
	err := pb.RegisterBookServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	return nil
}