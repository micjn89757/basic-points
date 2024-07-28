package main

import "inet/rpc/grpc/grpc-interpretor/server"

func main() {
	server := &server.Server{}
	server.Run()
}