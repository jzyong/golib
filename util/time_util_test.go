package util

import (
	"fmt"
	"testing"
	"time"
)

// 测试获取时间戳时间
func TestCurrentTimeSecond(t *testing.T) {
	SetTime("2022-02-23 16:00:00")
	second := CurrentTimeSecond()
	fmt.Printf("time Offset scond %v \n", second-time.Now().Unix())
}

// 测试是否同一周
func TestSameWeek(t *testing.T) {
	//2022-1-28	2022-1-30 true
	sameWeek := SameWeek(1643317200, 1643490000)
	fmt.Printf("same week:%v\n", sameWeek)

	//2022-1-28	2022-1-31 false false
	sameWeek = SameWeek(1643317200, 1643641624)
	fmt.Printf("same week:%v\n", sameWeek)
}

// 测试是当前时间是否在时间段内
func TestBetweenNow(t *testing.T) {
	between1 := BetweenNow("2022-09-01 00:00:00", "2023-09-01 00:00:00")
	fmt.Printf("between now:%v\n", between1)

	between2 := BetweenNow("2022-09-01 00:00:00", "2022-11-01 00:00:00")
	fmt.Printf("between now:%v\n", between2)
}
