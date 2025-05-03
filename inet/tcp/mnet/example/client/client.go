package main

import "inet/tcp"

func main() {
	client := tcp.NewClient("tcp6", ":3000")
	client.Run()
}