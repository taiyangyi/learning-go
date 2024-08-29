package main

import "fmt"

// 变量

// 局部全量声明
// 全局变量声明
// 多变量声明

// 三种声明方式
var a_1 int = 100
var a_2 = 100
var a_3 int

// := 该声明适用于函数体中

func main() {

	// 全局变量
	fmt.Println("a_1 = ", a_1, "a_2 = ", a_2, "a_3=", a_3) // a_1 =  100 a_2 =  100 a_3= 0

	// 方法一：声明一个变量，默认值为 0
	var a int
	fmt.Printf("type of a =%T\n", a)
	fmt.Println("value of a =", a)

	// 	type of a =int
	// value of a = 0

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T, value of b = %d", b, b) // type of b = int, value of b = 100

	var bb string = "abcd"
	fmt.Println("bb = ", bb)
	fmt.Printf("type of bb = %T, value of b = %s", bb, bb) // type of bb = string, value of b = abcd

	// 方法三：在初始化的时候，可以省去数据类型，通过自动匹配当前的变量类型
	// Go 将自动推断已经初始化的变量类型。
	var c = 188
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Println("cc = ", cc)                         // cc =  abcd
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc) // cc = abcd, type of cc = string

	// 方法四：(常用方法)省去var关键字，直接自动匹配
	d := 120
	fmt.Println("d = ", d)            // d =  120
	fmt.Printf("type of d = %T\n", d) // type of d = int

	// 一次性声明多个变量。
	var e, f, g = 1, 2, 3
	fmt.Printf("e = %d, f = %d, g = %d\n", e, f, g) // e = 1, f = 2, g = 3

	var xx, yy = 100, "aksk"
	fmt.Println("xx=", xx, "yy=", yy)                       // xx= 100 yy= aksk
	fmt.Printf("type of xx=%T and type of yy=%T\n", xx, yy) // type of xx=int and type of yy=string

	// 多行的多变量声明
	var (
		w int    = 100
		v bool   = true
		s string = "hello"
	)

	fmt.Printf("w = %d,v = %v,s = %s", w, v, s) // w = 100,v = true,s = hello

}
