package main

import "inet/rpc/grpc/grpc-validate/server"

func main() {
	server := &server.Server{}
	server.Run()
}