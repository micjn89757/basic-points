package commonutils

import (
	"fmt"
	"net/http"
)


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 沙河！")

}
// SeverDemo 默认的Server示例
func SeverDemo() {
	http.HandleFunc("/", helloWorld) // handleFunc可以向DefaultServeMux添加处理器
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}
	
}
