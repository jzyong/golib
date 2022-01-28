package util

import (
	"fmt"
	"testing"
)

//测试是否同一周
func TestSameWeek(t *testing.T) {
	//2022-1-28	2022-1-30
	sameWeek := SameWeekNow(1643555224)
	fmt.Printf("same week:%v\n", sameWeek)

	//2022-1-28	2022-1-31
	sameWeek = SameWeekNow(1643641624)
	fmt.Printf("same week:%v\n", sameWeek)
}
