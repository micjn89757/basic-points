package json

import (
	"time"
	"fmt"
)

type MyDate time.Time

var MyDateFormat = "2016-01-02"

// 序列化
func (d MyDate) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(d).Format(MyDateFormat))
	return []byte(s), nil
}


// 反序列化
func (d *MyDate) UnmarshalJSON(bs []byte) error {
	now, err := time.ParseInLocation(`"` + MyDateFormat + `"`, string(bs), time.Local)
	if err != nil {
		return err
	}

	*d = MyDate(now)

	return nil 
}


// Print(MyDate) 会调用String方法
func (d MyDate) String() string {
	return time.Time(d).Format(MyDateFormat)
}