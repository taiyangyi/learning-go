package main

import "fmt"

func main() {

	// 多维数组
	var arr [3][3]int // 定义一个3行3列的二维数组
	// 数组元素初始化
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = i*3 + j + 1
		}
	}

	// 输出数组元素
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("第 %d 行，第 %d 列的值 = %d\n", i+1, j+1, arr[i][j])
		}
	}

	// 第 1 行，第 1 列的值 = 1
	// 第 1 行，第 2 列的值 = 2
	// 第 1 行，第 3 列的值 = 3
	// 第 2 行，第 1 列的值 = 4
	// 第 2 行，第 2 列的值 = 5
	// 第 2 行，第 3 列的值 = 6
	// 第 3 行，第 1 列的值 = 7
	// 第 3 行，第 2 列的值 = 8
	// 第 3 行，第 3 列的值 = 9
}
