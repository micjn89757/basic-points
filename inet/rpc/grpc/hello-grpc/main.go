package main

import "inet/rpc/grpc/hello-grpc/server"

func main() {
	server := &server.Server{}
	server.Run()
}