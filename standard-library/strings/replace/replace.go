package main

import (
	"fmt"
	"strings"
)

// strings.replace(s, old, new, n) string
// 将字符串 s 前 n 个不重叠的 old 子串 都替换为 new 的新字符串
// 如果n<0会替换所有old子串。

// strings.replaceAll(s, old, new) string
// 将字符串s中的old子串全部替换为new的新字符串
func main() {

	str := "You are a good boy"
	newStr := strings.Replace(str, "good", "bad", 1)
	fmt.Println(newStr) // You are a bad boy

	str2 := "You are a good good good good boy"
	newStr2 := strings.Replace(str2, "good", "bad", 1)
	fmt.Println(newStr2) // You are a bad good good good boy

	newStr3 := strings.Replace(str2, "good", "bad", 2)
	fmt.Println(newStr3) // You are a bad bad good good boy

	newStr4 := strings.Replace(str2, "good", "bad", -1)
	fmt.Println(newStr4) // You are a bad bad bad bad boy

	////////////////ReplaceAll/////////////////

	str5 := "I like your small house"
	newStr5 := strings.ReplaceAll(str5, "house", "room")
	fmt.Println(newStr5) // I like your small room

	str6 := "I like your small small small small house"
	newStr6 := strings.ReplaceAll(str6, "small", "big")
	fmt.Println(newStr6) // I like your big big big big house

	// 按理，对一些敏感词替换
	str7 := "我曹，我曹，这狗东西"
	newStr7 := strings.Replace(str7, "狗", "*", 1)
	fmt.Println(newStr7) // 我曹，我曹，这*东西
	newStr8 := strings.ReplaceAll(newStr7, "曹", "*")
	fmt.Println(newStr8) // 我*，我*，这*东西
}
