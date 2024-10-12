package main

import "fmt"


// 指针地址
func main() {
	m := "hello"
	fmt.Printf("指针变量的值为 %p\n", &m) // 指针变量的值为 0xc000026070

	var mp *string         // 字符串指针变量
	mp = &m                // 通过变量取地址
	fmt.Printf("%p\n", mp) // 0xc000026070

	// 通过指针访问值
	fmt.Println(*mp)

}

// go run .\pointer.go
// 指针变量的值为 0xc000026070
// 0xc000026070
// hello
