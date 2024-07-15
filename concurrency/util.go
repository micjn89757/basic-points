package concurrency

import (
	"bytes"
	"encoding/binary"
)

// TODO: 整形转换成字节
func Int2Bytes(n int) []byte {
	x := int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}


// TODO: 字节转换成整型
func Bytes2Int(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64 
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}