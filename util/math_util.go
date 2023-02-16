package util

import (
	"github.com/jzyong/golib/log"
	"math/rand"
)

// RandomInt32 随机 int32 ，包含开始和结束
func RandomInt32(start, end int32) int32 {
	n := end - start
	if n < 0 {
		return start
	}
	return rand.Int31n(n+1) + start
}

// RandomInt64 随机 int64 ，包含开始和结束
func RandomInt64(start, end int64) int64 {
	n := end - start
	if n < 0 {
		return start
	}
	return rand.Int63n(n+1) + start
}

// RandomBool 根据几率 计算是否生成，种子数为10000
func RandomBool(probability int32) bool {
	randomSeed := rand.Int31n(10001)
	return probability >= randomSeed
}

// WightRandomTwo 根据权重随机,第二参数为权重
func WightRandomTwo[A any](drops []*Two[A, int32]) A {
	var a A
	if drops == nil {
		return a
	}
	var totalWeight int32
	for _, drop := range drops {
		totalWeight += drop.B
	}
	wight := RandomInt32(0, totalWeight)
	totalWeight = 0
	for _, drop := range drops {
		totalWeight += drop.B
		if wight <= totalWeight {
			return drop.A
		}
	}
	log.Warn("对象：%v ，随机数%v 未随机到对象？", drops, wight)
	return a
}
