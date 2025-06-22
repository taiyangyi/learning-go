package main

import "fmt"

func main() {
	// 声明局部变量，a 和 b 并赋值
	var a int = 5
	var b int = 6
	c := a + b
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
}

// a=5,b=6,c=11
