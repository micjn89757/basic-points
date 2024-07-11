package concurrency

import "testing"

func TestTraveseChannel(t *testing.T) {
	traveseChannel()
}

func TestNioChannel(t *testing.T) {
	nioChannel()
}

func TestBioChannel(t *testing.T) {
	bioChannel()
}

func TestSignalChannel(t *testing.T) {
	signalChannel()
}

