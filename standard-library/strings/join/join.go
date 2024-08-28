package main

import (
	"fmt"
	"strings"
)

func main() {

	// 	Join(elems []string, sep string) string
	// 	将字符串切片中的元素以指定字符串连接生成新字符串
	// 	第一个参数 elems 为字符串切片。
	// 第二个参数 sep 为连接符。
	// 返回值为新的字符串。

	strSlice := []string{"A", "tea", "or", "a", "coffer"}
	joinStrSlice := strings.Join(strSlice, "=")
	joinStrSlice2 := strings.Join(strSlice, " ")

	fmt.Println(joinStrSlice)  // A=tea=or=a=coffer
	fmt.Println(joinStrSlice2) // A tea or a coffer

}
