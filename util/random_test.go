package util

import (
	"fmt"
	"testing"
)

// 随机
func TestRandomInt32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v \n", RandomInt32(1000, 5000))
	}
}

// 随机bool
func TestRandomBool(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v \n", RandomBool(3000))
	}
}

// 根据权重随机,第二参数为权重
func TestWightRandomTwo(t *testing.T) {
	drops := []*Two[int32, int32]{{1, 7000}, {2, 3000}}
	one := 0
	two := 0
	for i := 0; i < 10000; i++ {
		result := WightRandomTwo[int32](drops)
		if result == 1 {
			one++
		} else {
			two++
		}
		//fmt.Printf("%v \n", result)
	}
	fmt.Printf("1=%v 2=%v \n", one, two)
}
