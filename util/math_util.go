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

//根据几率 计算是否生成，种子数为10000
func RandomBool(probability int32) bool {
	randomSeed := rand.Int31n(10001)
	return probability >= randomSeed
}
