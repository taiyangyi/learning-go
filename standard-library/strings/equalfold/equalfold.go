package main

import (
	"fmt"
	"strings"
)

func main() {

	// 区分大小写比较

	// 直接通过 == 操作符进行区分大小写的字符串比较即可
	str := "hello"
	str2 := "hello"
	str3 := "Hello"

	fmt.Println(str == str2) // true
	fmt.Println(str == str3) // false

	// 不区分大小写比较

	// 使用 EqualFold(s, t string) bool 函数进行比较，
	// 两个参数为需要比较的两个字符串，返回值为布尔值，
	// 如果是 true 说明字符串相等，反之 false 则字符串不相等。

	isEqual := strings.EqualFold(str, str2)
	isEqual2 := strings.EqualFold(str, str3)

	fmt.Println(isEqual)  // true
	fmt.Println(isEqual2) // true

}
