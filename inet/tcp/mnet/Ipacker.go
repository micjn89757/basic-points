package tcp

import (
	"io"
)

type IPacker interface {
	Pack(string) ([]byte, error) // 对tcp数据段编码
	Unpack(reader io.Reader) (string, error)        // 解码
}
