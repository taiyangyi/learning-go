package main

import (
	"fmt"
)

// `const` 用于声明一个常量。
const s string = "hello"

func main() {

	fmt.Println(s)

	//`const` 语句可以出现在任何 `var` 语句可以出现的地方
	const n = 5000

	// 常量表达式可以执行任意精度的运算
	const d = 10 / n
	fmt.Printf("Type of d is %T and value of d is %v\n", d, d) // Type of d is int and value of d is 0
}
