package main

import "inet/tcp"

func main() {
	server := tcp.NewServer(tcp.WithTCPAddr("tcp6", ":3000"))

	server.Run()
}