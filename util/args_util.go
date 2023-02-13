package util

import "strings"

// Two 两个参数结构体
type Two[A any, B any] struct {
	A A
	B B
}

// Three 三个参数结构体
type Three[A any, B any, C any] struct {
	A A
	B B
	C C
}

// Four 四个参数结构体
type Four[A any, B any, C any, D any] struct {
	A A
	B B
	C C
	D D
}

// Five 五个个参数结构体
type Five[A any, B any, C any, D any, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// ParseTwoArgs  解析字符串道具 splitSlice数组分割  splitArgs 参数分割
// 例如2=10,5=25,10=20,15=10,20=5 splitSlice
func ParseTwoArgs[A any, B any](configString, splitSlice, splitArgs string, parseFun func(configStr []string) (A, B)) []*Two[A, B] {
	itemStrs := strings.Split(configString, splitSlice)
	items := make([]*Two[A, B], 0, len(itemStrs))
	for _, str := range itemStrs {
		itemStr := strings.Split(str, splitArgs)
		a, b := parseFun(itemStr)
		items = append(items, &Two[A, B]{A: a, B: b})
	}
	return items
}

// ParseThreeArgs  解析字符串道具 splitSlice数组分割  splitArgs 参数分割
// 例如2=10=3,5=25=3,10=20=3,15=10=3,20=5=3 splitSlice
func ParseThreeArgs[A any, B any, C any](configString, splitSlice, splitArgs string, parseFun func(configStr []string) (A, B, C)) []*Three[A, B, C] {
	itemStrs := strings.Split(configString, splitSlice)
	items := make([]*Three[A, B, C], 0, len(itemStrs))
	for _, str := range itemStrs {
		itemStr := strings.Split(str, splitArgs)
		a, b, c := parseFun(itemStr)
		items = append(items, &Three[A, B, C]{A: a, B: b, C: c})
	}
	return items
}

// ParseFourArgs  解析字符串道具 splitSlice数组分割  splitArgs 参数分割
// 例如2=10=1=3,5=25=1=3,10=20=1=3,15=10=1=3,20=5=1=3 splitSlice
func ParseFourArgs[A any, B any, C any, D any](configString, splitSlice, splitArgs string, parseFun func(configStr []string) (A, B, C, D)) []*Four[A, B, C, D] {
	itemStrs := strings.Split(configString, splitSlice)
	items := make([]*Four[A, B, C, D], 0, len(itemStrs))
	for _, str := range itemStrs {
		itemStr := strings.Split(str, splitArgs)
		a, b, c, d := parseFun(itemStr)
		items = append(items, &Four[A, B, C, D]{A: a, B: b, C: c, D: d})
	}
	return items
}

// ParseFiveArgs  解析字符串道具 splitSlice数组分割  splitArgs 参数分割
// 例如2=10=1=3=1,5=25=1=3=1,10=20=1=3=1,15=10=1=3=1,20=5=1=3=1 splitSlice
func ParseFiveArgs[A any, B any, C any, D any, E any](configString, splitSlice, splitArgs string, parseFun func(configStr []string) (A, B, C, D, E)) []*Five[A, B, C, D, E] {
	itemStrs := strings.Split(configString, splitSlice)
	items := make([]*Five[A, B, C, D, E], 0, len(itemStrs))
	for _, str := range itemStrs {
		itemStr := strings.Split(str, splitArgs)
		a, b, c, d, e := parseFun(itemStr)
		items = append(items, &Five[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e})
	}
	return items
}
