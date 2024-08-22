package main

import (
	"fmt"
	"strings"
)

// string.split(s, sep string) []string
// 分隔字符串函数，返回一个字符串切片
// s  需要分隔的字符串
// sep 分隔标识
// 返回字符串切片，保留被分割出来的子字符串
func main() {

	str := "A-tea-or-a-coffer"
	strSlice := strings.Split(str, "-")
	fmt.Println(strSlice) // [A tea or a coffer]

}
