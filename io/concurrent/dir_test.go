package concurrent

import (
	"io_/util"
	"fmt"
	"testing"
	"time"
)

func TestProcessDir(t *testing.T) {
	dir := util.RootPath + "data/search_log" //日志文件存放的路径，file是相对于执行go run的路径
	files := ListDir(dir)
	begin := time.Now()
	ProcessDir(files)
	fmt.Printf("ProcessDir sum=%d, time %dms\n", sum, time.Since(begin).Milliseconds()) //9.8秒
}

func TestConcurrentProcessDir(t *testing.T) {
	dir := util.RootPath + "data/search_log" //日志文件存放的路径，file是相对于执行go run的路径
	files := ListDir(dir)
	begin := time.Now()
	ConcurrentProcessDir(files)
	fmt.Printf("ConcurrentProcessDir sum=%d, time %dms\n", sum, time.Since(begin).Milliseconds()) //4秒
}

//go test -v .\io\concurrent\ -run=^TestProcessDir$ -count=1
//go test -v .\io\concurrent\ -run=^TestConcurrentProcessDir$ -count=1
