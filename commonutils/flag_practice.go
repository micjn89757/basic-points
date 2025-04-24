package commonutils

import (
	"flag"
	"log"
	"time"
)


func Run() {
	// 定义命令行flag参数
	// flag.TypeVar(Type指针，flag名，默认值，帮助信息)
	var name string 
	var age int 
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	// 定义命令行flag参数的另一种方式
	// flag.Type(flag名，默认值，帮助信息)
	// 例如 name := flag.String("name", "张三", "姓名")


	// 定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析
	flag.Parse() 
	// 支持的命令行参数格式：
	// --flagname xxx/-flagname xxx/-flagname=xxx/--flag=xxx
	// flag在解析第一个非flag参数之前停止，或者在终止符"-"之后停止
	log.Println(name, age, married, delay)

	// 返回命令行参数后的其他参数, []stirng类型
	log.Println(flag.Args())

	// 返回命令行参数后的其他参数个数
	log.Println(flag.NArg())

	// 返回使用的命令行参数个数
	log.Println(flag.NFlag())
}
