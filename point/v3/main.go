package main

import "fmt"

func main() {

	// 关键字 * 表示指针调用符
	// 将 * 放在一个指针变量前，就会获得该指针变量对应的变量值
	x := 1024

	// y 变量表示一个指针变量，值对应着 x 的内存地址
	y := &x
	fmt.Printf("x 的内存地址%p\n", y) // x 的内存地址0xc00000a100

	// *y 表示 y 对应的变量的值，也就是 x 的值 1024
	fmt.Println("*y的值", *y) // *y的值 1024

	*y = 2121
	fmt.Println("修改*y的值,新值为", *y) // 修改*y的值,新值为 2121

	fmt.Println("x的值", x) // 那么对应的 x 的值也将被修改， x的值 2121

}
