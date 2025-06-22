package main

import "fmt"

type AInterface interface {
	test01()
}

type BInterface interface {
	test02()
}

type CInterface interface {
	AInterface
	BInterface
	test03()
}

type Stu struct {
}

func (s Stu) test01() {
	fmt.Println("test01")
}

func (s Stu) test02() {
	fmt.Println("test02")
}

func (s Stu) test03() {
	fmt.Println("test03")
}

func main() {

	var s Stu

	var c CInterface = s
	c.test01()
	c.test02()
	c.test03()
}
