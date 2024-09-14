package main

import "fmt"

// for 循环基本结构

// for 初始化变量; 条件判断; 修正变量 {
// 	循环体
// }

func main() {

	// 最基础的方式，单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 	1
	// 2
	// 3

	// 经典的初始/条件/后续 for 循环
	for j := 0; j < 9; j++ {
		if j == 8 {
			fmt.Println(j)
			break
		}
		fmt.Print(j)

	}
	// 012345678

	// 不带条件的 `for` 循环将一直重复执行，直到在循环体内使用
	// 了 `break` 或者 `return` 来跳出循环。
	for {
		fmt.Println("loop")
		break
	}

	// 使用 continue 跳过本次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// 1
	// 3
	// 5

}
