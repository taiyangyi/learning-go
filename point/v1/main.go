package main

import "fmt"

func swap(pa *int, pb *int) {
	temp := *pa
	*pa = *pb
	*pb = temp
}

func main() {

	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a = %v,b = %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("交换前 a = %v,b = %v\n", a, b)

}

// 输出：
// 交换前 a = 100,b = 200
// 交换前 a = 200,b = 100
