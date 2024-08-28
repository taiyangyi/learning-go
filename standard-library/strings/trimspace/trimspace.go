package main

import (
	"fmt"
	"strings"
)

func main() {

	str := " This is my English teacher    "
	trimSpaceStr := strings.TrimSpace(str)
	fmt.Println("原字符串:", str)
	fmt.Println("去掉左右空格的字符串:", trimSpaceStr)
}

// 原字符串  This is my English teacher
// 去掉左右空格的字符串 This is my English teacher
