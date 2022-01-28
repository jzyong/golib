package util

import "time"

//是否为同一天
func SameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

//是否同一周 单位为s
func SameWeek(a, b int64) bool {
	y1, w1 := time.Unix(a, 0).ISOWeek()
	y2, w2 := time.Unix(b, 0).ISOWeek()
	return y1 == y2 && w1 == w2
}

//和当前时间是否为同一周  单位为s
func SameWeekNow(a int64) bool {
	return SameWeek(time.Now().Unix(), a)
}

//是否为今天
func IsToday(ts int64) bool {
	return SameDay(time.Unix(ts, 0), time.Now())
}

func CheckTimeFormat(src, layout string) bool {
	_, err := time.Parse(layout, src)
	return err == nil
}

//获取零点时间
func ZeroUnixTime(offsetDay int) int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()*1000 + int64(offsetDay*24*3600*1000)
}
