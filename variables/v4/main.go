package main

import "fmt"

func main() {

	// 第一种写法
	var a = 1
	var b = 2
	var c int

	c = b
	b = a
	a = c
	fmt.Println("a", a)
	fmt.Println("b", b)

	//a 2
	//b 1

	// 第二种写法
	var m = 10
	var n = 20
	m = m ^ n
	n = n ^ m
	m = m ^ n
	fmt.Println("m", m)
	fmt.Println("n", n)
	//m 20
	//n 10

	// 第三种
	var w = 100
	var s = 200
	w, s = s, w
	fmt.Println("w", w)
	fmt.Println("s", s)

	//w 200
	//s 100

}
