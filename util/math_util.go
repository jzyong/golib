package util

import "math/rand"

//随机 int32 ，包含开始和结束
func RandomInt32(start, end int32) int32 {
	n := end - start
	if n < 0 {
		return start
	}
	return rand.Int31n(n+1) + start
}
