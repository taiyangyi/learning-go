package main

import "fmt"

// const 定义枚举类型
const (

	// 在const() 添加一个关键字iota,每行的iota都会累加1,第一行的iota的默认值为0
	SUCCESS = iota // 0
	FAILED         // 1
	OTHERS         // 2
)

// 注意：iota 只能配合const()一起使用，iota在const进行累加的时候使用

const (
	a, b = iota + 1, iota + 2  // iota=0; a=1;b=2
	c, d                       // iota=1; c=2;d=3
	e, f                       // iota=2; e=3;f=4
	g, h = iota * 2, iota * 10 // iota=3; g=6;h=30
	i, j                       // iota=4; i=8;j=40
)

func main() {

	fmt.Println("success=", SUCCESS)
	fmt.Println("failed=", FAILED)
	fmt.Println("others=", OTHERS)
	// 	success= 0
	// failed= 1
	// others= 2

	fmt.Println("a=", a, "b=", b)
	fmt.Println("c=", c, "d=", d)
	fmt.Println("e=", e, "f=", f)
	fmt.Println("g=", g, "h=", h)
	fmt.Println("i=", i, "j=", j)

}
