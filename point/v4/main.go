package main

import "fmt"

func main() {
	ok := true
	var p *bool            // 布尔型指针变量
	p = &ok                // 获取 ok 的地址
	fmt.Printf("%t\n", *p) // 输出指针变量p 对应的 ok 的值  // true

	*p = false
	fmt.Printf("%t\n", *p) // false
	fmt.Printf("%t\n", ok) // false
}
