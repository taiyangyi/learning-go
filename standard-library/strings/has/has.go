package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "0xhello001x0"

	// 	HasPrefix(s, prefix string) bool
	// //判断字符串是否以指定的字符串开头
	// 第一个参数 s 为被判断字符串。
	// 第二个参数 prefix 为指定的字符串。

	flagHasPrefix := strings.HasPrefix(str, "0x")
	fmt.Println(flagHasPrefix) // true

	// 	HasSuffix(s, prefix string) bool
	// 	判断字符串是否以指定的字符串结束
	// 	第一个参数 s 为被判断字符串。
	// 第二个参数 prefix 为指定的字符串。

	flagHasSuffix := strings.HasSuffix(str, "x0")
	fmt.Println(flagHasSuffix) // true

	flagHasSuffix_2 := strings.HasSuffix(str, "hello")
	fmt.Println(flagHasSuffix_2) // false

}
