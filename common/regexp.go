package commonutils

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// createRegExpObj 创建正则表达式对象
func createRegExpObj() (bool, error) {
	str := "Hello, world!"
	// 创建正则表达式对象
	// `[aeiou]` 这个叫做正则表达式字面量
	re, err := regexp.Compile(`[aeiou]`)

	// 这种方式如果创建失败会直接panic
	// regexp.MustCompile()
	if err != nil {
		return false, fmt.Errorf("compile regexp object failed, err: %w", err)
	}

	log.Println(re.MatchString(str))

	return re.MatchString(str), nil
}

// matchStringBytes 字符串和字节切片匹配
func matchStringBytes() (bool, error) {
	// 字符串使用MatchString匹配
	// 字节切片使用Match匹配
	str := []byte("helloworld")

	re, err := regexp.Compile(`^Hello`)

	if err != nil {
		return false, err
	}

	return re.Match(str), nil 

}

// strReplace 字符串替换
func strReplace() string {
	str := "hello world!"

	// 编译正则表达式对象
	re := regexp.MustCompile(`world`)

	// 替换所有world
	newStr := re.ReplaceAllString(str, "golang")
	
	// ReplaceAll传入的是[]byte

	// 动态替换内容
	re2 := regexp.MustCompile(`golang`)
	resStr := re2.ReplaceAllStringFunc(newStr, strings.ToUpper)
	return resStr
}

// findSubMatch 捕获组
func findSubMatch() []string {
	/// 捕获组是用括号括起来的子表达式，可以提取特定的子字符串
	str := "John Doe, jane@example.com"
	re := regexp.MustCompile(`(\w+)\s(\w+),\s(\w+@\w+.\w+)`)

	match := re.FindStringSubmatch(str)
	// FindAllStringSubmatch(s string, n int): 可以用n控制匹配的数量

	// 返回的[]string中第一个元素是整个表达式匹配到的字符串
	// 后面几个元素分别是满足括号中正则表达式的子字符串


	return match
}
