package main

import "fmt"

// 单个返回值
// x,y 为形参
// 返回值 int
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}

// 多个返回值，不指定名称
func getNumber(x, y, z int) (int, int, int) {
	return x, y, z
}

func main() {
	// 单个返回值
	max := max(8, 9)
	fmt.Printf("max = %d\n", max) // max = 9

	// 多个返回值，不指定名称
	x, y, z := getNumber(10, 11, 12)
	fmt.Printf("x=%d,y=%d,z=%d", x, y, z)

}
