package tcp

import (
	"encoding/binary"
	"fmt"
	"net"
)

var packer = &Packer{binary.BigEndian}

type server struct {
	listener net.Listener
	packer   IPacker
}

type ServerOptions func(*server)

func WithTCPAddr(network, addr string) ServerOptions {
	tcpAddr, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		panic(err)
	}
	return func(s *server) {
		s.listener, err = net.ListenTCP(network, tcpAddr)

		if err != nil {
			panic(err)
		}
	}
}

func WithPakcer(byteOrder binary.ByteOrder) ServerOptions {
	return func(s *server) {
		s.packer = &Packer{ByteOrder: byteOrder}
	}
}

func NewServer(options ...ServerOptions) *server {
	serv := &server{}
	if len(options) > 0 {
		for _, opt := range options {
			opt(serv)
		}
	} else {
		tcpOpt := WithTCPAddr("tcp6", ":8080")
		tcpOpt(serv)
	}
	
	return serv
}


func (s *server) Run() {
	// 等待客户连接
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if err = err.(net.Error);  err != nil {
				continue
			}
		}


		// for {
			// 解包
			msg, err := packer.Unpack(conn)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(msg)
		// }
		
	}
}