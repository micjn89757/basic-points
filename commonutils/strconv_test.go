package commonutils

import (
	"log"
	"os"
	"strconv"
	"testing"
)

// strconv实现了基本数据类型和其字符串表示的相互转换
func TestStrconv(t *testing.T) {
	logger := log.New(os.Stdout, "[strconv]", log.Lshortfile | log.Ldate)
	//Atoi 将字符串类型整数转换为int类型, 这一组用的最多
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		logger.Println("can't convert to int")
	}else {
		logger.Printf("type:%T value:%#v\n", i1, i1)
	}
	// Itoa 将int类型数据转换成对应的字符串表示
	i2 := 200
	s2 := strconv.Itoa(i2)
	logger.Printf("type:%T, value:%#v\n", s2, s2)


	// 将字符串转换为其他给定的值
	// ParseBool, ParseInt, ParseUnit, ParseFloat

	// 将数值转换成字符串
	// FormatBool, FormatInt, FormatUnit, FormatUint, FormatFloat
}