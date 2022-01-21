package util

import "time"

//是否为同一天
func SameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

func IsToday(ts int64) bool {
	return SameDay(time.Unix(ts, 0), time.Now())
}

func CheckTimeFormat(src, layout string) bool {
	_, err := time.Parse(layout, src)
	return err == nil
}
