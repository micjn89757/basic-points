package main

import (
	"inet/httpapi"
	"log"
)

func main() {
	server := httpapi.NewAPIServer(":8080")
	err := server.Run()
	if err != nil {
		log.Println(err)
	}
}