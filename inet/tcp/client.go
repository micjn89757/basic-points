package tcp

import (
	"encoding/binary"
	"net"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// var packer = &tcp.Packer{ByteOrder: binary.BigEndian}

type client struct {
	network	string
	addr  	string
	packer 	IPacker
}

func NewClient(network, addr string) *client {
	return &client{
		network: network,
		addr: addr,
		packer: &Packer{ByteOrder: binary.BigEndian},
	}
}


func (c *client) Run() {
	conn, err := net.Dial(c.network, c.addr)

	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go c.write(conn)
	wg.Wait()
}


func (c *client) write(conn net.Conn) {
	defer wg.Done()
	// for {
		conn.(*net.TCPConn).SetWriteDeadline(time.Now().Add(time.Second * 10))
		byt, err := c.packer.Pack("hello world")
		if err != nil {
			panic(err)
		}
		_, err = conn.Write(byt)
		if err != nil {
			panic(err)
		}
	// }
}