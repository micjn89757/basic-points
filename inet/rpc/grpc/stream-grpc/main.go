package main 

import "inet/rpc/grpc/stream-grpc/server"

func main() {
	serv := &server.Server{}

	serv.Run()
}