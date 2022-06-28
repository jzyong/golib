package util

import (
	"fmt"
	"testing"
)

//随机
func TestRandomInt32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v \n", RandomInt32(1000, 5000))
	}
}

//随机bool
func TestRandomBool(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v \n", RandomBool(3000))
	}
}
