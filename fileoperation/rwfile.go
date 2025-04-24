package fileoperation

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

// GetCurrentPath 获取当前执行程序路径
func GetCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}

	return "./"
}

// ReadFileBufio 带有缓冲区的读文件方式(常用)
func ReadFileBufio(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		log.Println("open file error:", err)
		return err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var resStr string

	for {
		line, err := reader.ReadString('\n') // 读到的结果包含\n

		if err != nil {
			if err == io.EOF {
				log.Println("read end")
				break
			} else {
				log.Println("file read err", err)
				return err
			}
		}

		resStr += line
	}

	log.Print("file content:\n", resStr)

	return nil
}

// WriteFileBufio 带缓冲区向文件写入内容
func WriteFileBufio(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("read file error:", err)
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		n, err := writer.WriteString("asdasd\n") // 将文件写入缓存
		if err != nil {
			log.Println("writer file error", err)
			return err
		}
		log.Printf("write %d bytes to file", n)
	}
	writer.Flush() // 刷新缓存

	return nil
}

// ReadFileBase 使用默认无缓冲的方式进行循环读取
func ReadFileBase(filename string) error {
	file, err := os.Open(filename) // 底层调用的是os.OpenFile(filename, os.RDONLY, 0)
	if err != nil {
		log.Println("文件打开失败", err)
		return err
	}

	defer file.Close()

	content := make([]byte, 100)
	res := make([]byte, 128)
	for {
		n, err := file.Read(content)
		if err == io.EOF { // 读到文件末尾会返回io.EOF
			log.Println("文件读完了")
			break
		}
		if err != nil {
			log.Println("read file error", err)
			return err
		}

		res = append(res, content[:n]...)

	}

	log.Println("文件内容:", string(res))

	return nil
}

// WriteFileBase 无缓冲写入文件基本使用
func WriteFileBase(filename string, wcontent string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("文件打开失败", err)
		return err
	}

	defer file.Close()

	n, err := file.Write([]byte(wcontent))

	if err != nil {
		log.Println("文件写入失败", err)
		return err
	}

	log.Printf("写入%d字节数据", n)

	return nil
}
