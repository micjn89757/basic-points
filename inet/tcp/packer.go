/*
防止粘包
*/

package tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type Packer struct {
	ByteOrder binary.ByteOrder	// 转化byte流的方式:大端还是小端，注意client和server和方式要一致
}

// 数据段长度记录 2B + 数据本身
func (p *Packer) Pack(msg string) ([]byte, error) {
	// 数据包长度转换成uint16  2B 
	len := uint16(len(msg) + 2)
	// 初始化要返回的数据包
	bf := make([]byte, 0)
	pkg := bytes.NewBuffer(bf)

	// 写入数据段长度记录
	err := binary.Write(pkg, p.ByteOrder, len)
	// fmt.Println(pkg.Len())	// 2

	if err != nil {
		return nil, err
	}

	// 写入数据
	err = binary.Write(pkg, p.ByteOrder, []byte(msg))
	// fmt.Println(pkg.Len())	// 13

	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}


func (p *Packer) Unpack(reader io.Reader) (string, error) {
	reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 10))
	bufReader := bufio.NewReader(reader)	// 转换成bufio
	lenBytes, err := bufReader.Peek(2)	// 前两个字节表示长度, 这里也不会真正从buffer中读取数据
	if err != nil {
		return "", err
	}

	fmt.Println(bufReader.Buffered())	// 13 => len + "hello world"

	lenBuff := bytes.NewBuffer(lenBytes)

	fmt.Println(len(lenBuff.Bytes()))	// 2

	var length uint16
	err = binary.Read(lenBuff, p.ByteOrder, &length)	// 把lenBuff中的数据读走

	fmt.Println(bufReader.Buffered())	// 13
	fmt.Println(len(lenBuff.Bytes()))	// 0
	if err != nil {
		return "", err 
	}

	// buffer中的现有的可读取字节数小于length
	if bufReader.Buffered() < int(length) {
		return "", err
	}

	// 读取消息数据
	pack := make([]byte, length)
	_, err = bufReader.Read(pack)	// bufReader.Reader会将数据读出
	if err != nil {
		return "", err
	}

	fmt.Println(bufReader.Buffered())

	return string(pack[2:]), nil
}