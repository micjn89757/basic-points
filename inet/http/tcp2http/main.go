package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"slices"
	"time"
)

type HttpServer struct {
	tcpListener *net.TCPListener
}


func NewServer(addr string) *HttpServer {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(err)
	}

	tcpListener, _ := net.ListenTCP("tcp4", resolveTCPAddr)
	s := &HttpServer{tcpListener: tcpListener}
	return s
}


func (s *HttpServer) Run() error {
	var tempDelay time.Duration 
	for {	// 不断监听连接
		conn, err := s.tcpListener.Accept()
		if err != nil {
			if err := err.(net.Error); err != nil && err.Timeout() {  // 如果是network error, 并且是因为网络阻塞导致的连接失败，等待一会再监听下一个连接
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}

				if max := 1 * time.Second; tempDelay > max { // 但是等待时长最大1s
					tempDelay = max 
				}

				slog.Info("http accept error", slog.String("err:",err.Error()))
				time.Sleep(tempDelay)
				continue
			}

			return err
		}


		// 一个协程处理
		go func() {
			for {	// 因为一个连接会接收多个http请求，所以要不断处理
				reader := bufio.NewReader(conn)
				req, _ := http.ReadRequest(reader)
				fmt.Println(req.URL, req.Host)
				conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nhello world"))
			}
		}()
	}
}

func main() {
	s := NewServer(":8080")
	s.Run()
}