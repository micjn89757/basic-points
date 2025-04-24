package times

import (
	"testing"
	"time"
)

func  TestTime(t *testing.T) {
	now := time.Now() //time.Time类型 获取现在时间

	// unix时间戳
	t.Logf("now: %v, nowUnix: %v",  now, now.Unix()) //秒级时间戳
	t.Logf("毫秒级别时间戳: %v, 微秒: %v, 纳秒: %v", now.UnixMilli(), now.UnixMicro(), now.UnixNano())

	// 时间戳转换为时间对象
	sec := now.Unix()
	timeObj := time.Unix(sec, 22) // 第二个参数为不足1s的纳秒数
	t.Logf("timeObj: %v", timeObj)
	// 其他时间戳转换一样, 但是只有纳秒级别转换必须要第二个参数

	// 获取time.Time对象年月日时分秒
	t.Logf("year:%v, month: %v, day: %v, hour: %v, min: %v, second: %v", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// time中提供了，time.Hour, time.Month等单位时间长度, 是time.Duration类型， 代表两个事件之间经过是时间，可表示最长时间段大约为290年，以纳秒为单位(time.Duration表示1ns)
	t.Logf("2hours: %v", 2 * time.Hour)

	
	// 时间操作
	// Time - Time = Duration Sub
	// Time + Duration = Time ADD
	// Time - Duration操作可以等价为Time + (-Duration) ADD
	later := now.Add(time.Hour)
	t.Logf("later: %v", later)
	// 比较时间是否相同, time的Equal方法，这个比较会考虑时区影响
	// 比较两个时间点先后顺序，Before, After

	
	
	// timezoneDemo 时区示例
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		t.Logf("load America/New_York location failed, err: %v", err)
		return
	}
	// 加载上海所在的时区
	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	t.Logf("timesAreEqual: %v",timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	t.Logf("timesAreEqual: %v",timesAreEqual)



	// formatDemo 时间格式化
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	t.Logf("%v", now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	t.Logf("%v", now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	t.Logf("%v", now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	t.Logf("%v", now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	t.Logf("%v", now.Format("15:04:05"))
	// 只格式化日期部分
	t.Logf("%v", now.Format("2006.01.02"))



	// 解析字符串格式为time.Time
	// parseDemo 指定时区解析时间
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj2, err := time.Parse("2006/01/02 15:04:05", "2022/10/05 11:25:20")
	if err != nil {
		return
	}
	t.Logf("%v", timeObj2)// 2022-10-05 11:25:20 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj2, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		return
	}
	// 2022-10-05 11:25:20 +0800 CST
	t.Logf("%v", timeObj2)

	// time.ParseInLocation 可以在解析时指定时区信息，指定时区信息使用time.LoacLocation/time.FixedZone

	t.Error("Test")
}