package main

import "fmt"

func main() {

	ok := true
	var p *bool // 布尔指针变量
	// var p *bool = &ok
	p = &ok                       // 获取 ok 对应的内存地址
	fmt.Println("ok对应的内存地址:", p)  // ok对应的内存地址: 0xc00000a0a8
	fmt.Printf("获取p的值: %t\n", *p) // 输出指针变量 p 对应的变量 ok 的值

	*p = false                     // 修改布尔指针变量对应的值
	fmt.Printf("获取修改后的值：%t\n", *p) // 输出指针变量 p 对应的变量 ok 的值 // 获取修改后的值：false

}

// go run .\pointer.go
// ok对应的内存地址: 0xc00000a0a8
// 获取p的值: true
// 获取修改后的值：false
