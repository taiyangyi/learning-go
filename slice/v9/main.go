package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 字符串转化为 []byte

	str := "helloworld"
	convbyte := []byte(str)
	fmt.Printf("convbyte type = %T, val = %s\n", convbyte, convbyte)
	// convbyte type = []uint8, val = helloworld

	// []byte 转换为字符串
	b := []byte{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'}
	s := string(b)
	fmt.Printf("s type = %T, val = %s\n", s, s)
	// s type = string, val = helloworld

	// 长度计算
	fmt.Printf("b length = %d\n", len(b))
	// b length = 10

	// 中文算作 3 个字符
	by := []byte("世界")
	fmt.Printf("by length = %d\n", len(by))
	// by length = 6

	// 中文算作 1 个字符
	fmt.Printf("by length = %d\n", utf8.RuneCount(by))
	// by length = 2

}
