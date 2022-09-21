package util

import (
	"fmt"
	"testing"
)

//获取运行APP路径
func TestGetAppPath(t *testing.T) {
	fmt.Printf("运行地址 %s", GetAppPath())
}
