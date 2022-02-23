package util

import (
	"fmt"
	"testing"
	"time"
)

//测试获取时间戳时间
func TestCurrentTimeSecond(t *testing.T) {
	SetTime("2022-02-23 16:00:00")
	second := CurrentTimeSecond()
	fmt.Printf("time Offset scond %v \n", second-time.Now().Unix())
}

//测试是否同一周
func TestSameWeek(t *testing.T) {
	//2022-1-28	2022-1-30 true
	sameWeek := SameWeek(1643317200, 1643490000)
	fmt.Printf("same week:%v\n", sameWeek)

	//2022-1-28	2022-1-31 false false
	sameWeek = SameWeek(1643317200, 1643641624)
	fmt.Printf("same week:%v\n", sameWeek)
}
