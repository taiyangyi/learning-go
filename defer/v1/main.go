package main

import "fmt"

// 知识点1：defer的执行顺序

func funcA() {
	fmt.Println("A")
}

func funcB() {
	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {

	// defer 的执行顺序（压出栈方式） c -> b -> a
	defer funcA()
	defer funcB()
	defer funcC()

}

// 输出结果：
// C
// B
// A
