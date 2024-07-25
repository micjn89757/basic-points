package json

import (
	"fmt"
	"time"
)

// 自定义json里的时间格式，核心是自定义一个type，实现MarshallJSON和UnmarshallJSON两个方法

var MyDateFromat = "2006-01-02"

type MyDate	time.Time


func (d MyDate) MarshallJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", time.Time(d).Format(MyDateFromat)) // 字符串里面加入""字符串是因为json字符串中的字符串还需要""
	return []byte(s), nil
}


// 需要改变自己，所以要传指针
func (d *MyDate) UnmarshallJSON(bs []byte) error {
	now, err := time.ParseInLocation(`"` + MyDateFromat + `"`, string(bs), time.Local) // 注意MyDateFormat前后还要加引号并且使用本机时区
	*d = MyDate(now)
	return err 
}


// print(MyDate)会调用String()方法
func (d MyDate) String() string {
	return time.Time(d).Format("2006-01-02")
}