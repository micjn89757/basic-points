package collections

import "fmt"

func MapDemo() map[string]string {
	// 初始化
	scoreMap := make(map[string]int, 8)

	scoreMap["a"] = 90
	scoreMap["b"] = 100

	// 判断某个key是否存在
	v, ok := scoreMap["a"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no")
	}

	// 方式2
	userInfo := map[string]string{
		"username": "shahe",
		"password": "123345",
	}

	return userInfo
}
