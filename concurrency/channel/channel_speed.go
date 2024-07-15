package channel

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// TODO: 有点问题，测试便知
// 多个文件合并
// 通过channel实现类似WaitGroup的操作
// 将channel作为多个协程的缓冲区，能够缓解上下游速度差


const PRODUC_NUM = 3

var buffer = make(chan string, 100)  // 行
var pc_sync = make(chan struct{}, PRODUC_NUM) // 信号
var all_over = make(chan struct{}) 


// 读一个文件，并把每行的内容放入channel
func producer(filename string) {
	fin, err := os.OpenFile(filename, os.O_RDONLY | os.O_CREATE, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer fin.Close()

	reader := bufio.NewReader(fin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF { // 表示读到最后一行
				if len(line) > 0 {
					buffer <- (line + "\n")
				}
				break
			} else {
				fmt.Println(err)
			}
		} else {
			buffer <- line
		}
		
	}

	<- pc_sync  // 表示生产结束
}

func consumer(filename string) {
	fout, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer fout.Close()

	writer := bufio.NewWriter(fout)

	for {
		if len(buffer) == 0 { // 生产者可能结束了，也可能是消费过快
			if len(pc_sync) == 0 { // 生产者都结束了
				break
			} else {
				time.Sleep(100 * time.Microsecond) // 减减速
			}
		} else {
			line := <- buffer
			writer.WriteString(line)
		}
	}

	writer.Flush()

	all_over <- struct{}{} // 表示消费完所有生产者的产品
}


func merge() {
	// 添加三个生产者
	for i := 0; i < PRODUC_NUM; i++ {
		pc_sync <- struct{}{}
	}

	go producer("1.txt")
	go producer("2.txt")
	go producer("3.txt")
	go consumer("merge.txt")

	<- all_over
}