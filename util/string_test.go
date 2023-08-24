package util

import (
	"fmt"
	"testing"
)

// 测试字节转换
func TestByteConvertString(t *testing.T) {
	fmt.Println(ByteConvertString(23))
	fmt.Println(ByteConvertString(123))
	fmt.Println(ByteConvertString(1230))
	fmt.Println(ByteConvertString(1234567))
	fmt.Println(ByteConvertString(123456789))
	fmt.Println(ByteConvertString(9123456789))
}

// 深拷贝
func TestDeepCopy(t *testing.T) {
	type Params struct {
		Id   int64
		Name string
	}
	src := &Params{
		Id:   1,
		Name: "golib",
	}
	dst := &Params{}
	DeepCopy(dst, src)
	fmt.Println(dst)

}
