package util

import (
	"fmt"
	"testing"
)

//计算字符串hash值 int
func TestGetIntHash(t *testing.T) {
	fmt.Printf("hash Id: %v \n", GetJavaIntHash("192.168.0.2"))
	fmt.Printf("hash Id: %v \n", GetJavaIntHash("192.168.0.1"))
	fmt.Printf("hash Id: %v \n", GetJavaIntHash("192.168.0.3"))
	fmt.Printf("hash Id: %v \n", GetJavaIntHash("192.168.110.3"))
	fmt.Printf("hash Id: %v \n", GetJavaIntHash("127.0.0.1"))
}
