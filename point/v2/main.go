package main

import "fmt"

func main() {

	// 取地址
	a := "helloworld"
	// 将 & 放到一个变量前，就会获得该变量对应的内存地址
	// p 变量是一个指针变量，值对应着变量 a 的地址
	p := &a
	fmt.Println("变量a的内存地址:", p) // 变量a的内存地址: 0xc0000262a0

	pi := 3.1415
	fmt.Printf("%p\n", &pi) // 直接获取pi变量对应的0xc00000a0d8

	var fp *float64 = &pi
	fmt.Printf("%p\n", fp) // 0xc00000a0d8
}
