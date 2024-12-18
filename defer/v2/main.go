package main

import "fmt"

// defer 和 return 谁先谁后

func deferFunc() int {
	fmt.Println("执行deferFunc()")
	return 0
}

func returnFunc() int {
	fmt.Println("执行returnFunc()")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {

	/**
	执行结果：
	执行returnFunc()
	执行deferFunc()

	可以看到 defer在return结束之后
	*/
	deferAndReturn()

}
