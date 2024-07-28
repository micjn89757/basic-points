package main

import "inet/rpc/grpc/grpc-token-auth/server"

func main() {
	server := &server.Server{}
	server.Run()
}