package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	Contains(s, substr string) bool
	// 	查找子串是否存在于指定的字符串中
	// 	第一个参数 s 为指定的字符串。
	// 第二个参数 substr 为子串。
	// 返回值为布尔值，如果是 true 说明存在，反之 false 则不存在。

	str := "I like summer and winter"
	strContains := strings.Contains(str, "summer")
	fmt.Println(strContains) // true
}
