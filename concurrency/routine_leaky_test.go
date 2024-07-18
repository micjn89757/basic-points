package concurrency

import "testing"

func TestClient(t *testing.T) {
	server()
	client()
}