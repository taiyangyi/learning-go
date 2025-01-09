package main

import "fmt"

func main() {
	num := []int{1, 2, 3}
	num2 := make([]int, 3)
	// copy前数据: [0 0 0]
	fmt.Printf("copy前数据: %v\n", num2)
	// func copy(dst, src []Type) int
	// 将 num 中的值依次拷贝到 num2 中
	copy(num2, num)
	// copy后数据: [1 2 3]
	fmt.Printf("copy后数据: %v\n", num2)
}
