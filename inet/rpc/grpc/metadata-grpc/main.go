package main

import "inet/rpc/grpc/metadata-grpc/server"

func main() {
	server := &server.Server{}
	server.Run()
}