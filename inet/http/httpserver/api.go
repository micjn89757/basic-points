package httpserver

import (
	"log"
	"net/http"
)

// !一个简单的http api server的模板示例

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	// 实现子路由
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	chain := MiddlewareChain(RequireAuthMiddleware, RequestLoggerMiddleware)

	server := &http.Server{
		Addr: s.addr,
		// Handler: RequireAuthMiddleware(RequestLoggerMiddleware(router)), // 要防止一层套一层的代码出现
		Handler: chain(v1),
	}

	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}

// 中间件类型
type Middleware func(http.Handler) http.HandlerFunc

// 方便注册多个中间件
func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		// 注意顺序
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}

// 简单中间件
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // 执行后面的逻辑
	}
}

// 简单身份验证中间件
func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		// check if the user is authenticated
		token := r.Header.Get("Authorization")
		if token != "Bearer Token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return 
		}

		next.ServeHTTP(w, r)
	}
}