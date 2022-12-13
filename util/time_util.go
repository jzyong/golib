package util

import (
	"github.com/jzyong/golib/log"
	"time"
)

//自定义系统时间，可以内部设置偏移量

// 时间偏移量 秒
var timeOffsetSecond int64 = 0

// 设置系统时间
func SetTime(timeStr string) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		log.Error("parse Time %v error:%v", timeStr, err)
	}
	timeOffsetSecond = t.Unix() - time.Now().Unix()
}

// 获取 服务器内部时间戳s，可能有偏移量
func CurrentTimeSecond() int64 {
	return time.Now().Unix() + timeOffsetSecond
}

// 获取 服务器内部时间戳 ms，可能有偏移量
func CurrentTimeMillisecond() int64 {
	return Now().UnixNano() / 1000000
}

// 当前时间，有时间偏移量
func Now() time.Time {
	return time.Unix(CurrentTimeSecond(), int64(time.Now().Nanosecond()))
}

// 是否为同一天
func SameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

// 是否同一周 单位为s
func SameWeek(a, b int64) bool {
	y1, w1 := time.Unix(a, 0).ISOWeek()
	y2, w2 := time.Unix(b, 0).ISOWeek()
	return y1 == y2 && w1 == w2
}

// 和当前时间是否为同一周  单位为s
func SameWeekNow(a int64) bool {
	return SameWeek(Now().Unix(), a)
}

// 是否为今天
func IsToday(ts int64) bool {
	return SameDay(time.Unix(ts, 0), Now())
}

// 检测时间格式是否正确
func CheckTimeFormat(src, layout string) bool {
	_, err := time.Parse(layout, src)
	return err == nil
}

// 获取零点时间
func ZeroUnixTime(offsetDay int) int64 {
	t := Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()*1000 + int64(offsetDay*24*3600*1000)
}

// BetweenNow 是否在当前时间内
func BetweenNow(startTime, endTime string) bool {
	start, err := time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		log.Error("parse start time %v error:%v", startTime, err)
		return false
	}
	end, err := time.Parse("2006-01-02 15:04:05", endTime)
	if err != nil {
		log.Error("parse end time %v error:%v", endTime, err)
		return false
	}
	now := Now()
	return now.Unix() > start.Unix() && now.Unix() < end.Unix()
}
