package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// ToStringIndent 将任意结构转化为json缩进后的字符串 方便输出调试
func ToStringIndent(what interface{}) string {
	b, err := json.Marshal(what)
	if err != nil {
		return fmt.Sprintf("%+v", what)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	return out.String()
}

// ToString 将任意结构转化为json字符串 方便输出调试
func ToString(what interface{}) string {
	b, err := json.Marshal(what)
	if err != nil {
		return fmt.Sprintf("%+v", what)
	}
	return string(b)
}

// 是否包含某个元素
func SliceContains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// 字符串转int64,失败返回0
func ParseInt64(numberStr string) int64 {
	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		log.Printf("%v 转换异常：%v", numberStr, err)
		return 0
	}
	return number
}

// 字符串转int32,失败返回0
func ParseInt32(numberStr string) int32 {
	number, err := strconv.ParseInt(numberStr, 10, 32)
	if err != nil {
		log.Printf("%v 转换异常：%v", numberStr, err)
		return 0
	}
	return int32(number)
}

// 字符串转int,失败返回0
func ParseInt(numberStr string) int {
	number, err := strconv.ParseInt(numberStr, 10, 32)
	if err != nil {
		log.Printf("%v 转换异常：%v", numberStr, err)
		return 0
	}
	return int(number)
}

// is the string a port
func IsPort(p string) bool {
	pi, err := strconv.Atoi(p)
	if err != nil {
		return false
	}
	if pi > 65536 || pi < 1 {
		return false
	}
	return true
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
