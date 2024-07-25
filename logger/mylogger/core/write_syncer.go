/*
	定义日志写入目标类型
*/
package core

import (
	"io"
	"sync"
)

// 例如os.stderr,os.stdout,os.File 都实现了这个接口
type WriteSyncer interface {
	io.Writer 
	Close() error
	Sync() error	// 例如内存数据刷新到磁盘
}


// 实现一个并发安全的WriteSyncer
type lockedWriteSyncer struct {
	sync.Mutex
	ws WriteSyncer
}

// 数据写入
func (ws *lockedWriteSyncer) Write(bs []byte) (int, error) {
	ws.Lock()
	defer ws.Unlock()
	n, err := ws.ws.Write(bs)
	return n, err
}

// 关闭句柄
func (ws *lockedWriteSyncer) Close() error {
	err := ws.ws.Close()
	return err
}

// 数据刷新到磁盘
func (ws *lockedWriteSyncer) Sync() error {
	ws.Lock()
	defer ws.Unlock()
	err := ws.ws.Sync()	
	return err 
}

// 对并发不安全的WriteSyncer进行包装，比如*os.Files在使用前必须要要lock，否则多个goroutine会对文件进行覆盖
func Lock(ws WriteSyncer) WriteSyncer {
	if _, ok := ws.(*lockedWriteSyncer); ok {
		return ws
	}
	return &lockedWriteSyncer{ws: ws}
}


// ? 这个地方io.Writer embed到了结构体中，相当于writerWrapper继承了io.Writer的具体实现对象
type writerWrapper struct {
	io.Writer
}

func (w writerWrapper) Sync() error {
	return nil
}


func (w writerWrapper) Close() error {
	return nil
}


// AddSync 将io.Writer转换为WriteSyncer，如果io.Writer的具体类型实现了WriterSyncer，就可以使用它的Sync()，如果没有我们会添加一个不进行任何操作的Sync()方法
func AddSync(w io.Writer) WriteSyncer {
	switch w := w.(type) {
	case WriteSyncer:
		return w
	default:
		return writerWrapper{w} // writerWrapper继承了w的Write方法
	}
}