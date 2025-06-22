package main

import "fmt"

type AInterface interface {
	doSomething()
}

type integer int

func (i integer) doSomething() {
	fmt.Println("只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型,i=", i)
}

func main() {

	var i integer = 10
	var a AInterface = i
	a.doSomething()
}
