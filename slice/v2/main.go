package main

import "fmt"

func main() {

	// 声明一个切片，并且初始化默认值是 1,2,3,4,5 长度为 5
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Printf("len=%d,s=%v\n", len(s), s)

	// 	[1 2 3 4 5]
	// len=5,s=[1 2 3 4 5]

	fmt.Println("=====================================")

	// 声明一个 s2 切片，同时给 s2 分配空间，长度为3，初始化值为0
	var s2 []int = make([]int, 3)
	fmt.Println(s2)
	// [0 0 0]

	fmt.Println("=====================================")

	// 声明s3是一个切片，同时分配空间，初始化值为0，通过:=推到出s3是一个切片
	s3 := make([]int, 4)
	fmt.Println(s3)
	// 判断切片是否存在元素
	if len(s3) == 0 {
		fmt.Println("s3 是一个空切片")
	} else {
		fmt.Println("s3 是一个非空切片")
	}
	// 	[0 0 0 0]
	// s3 是一个非空切片

}
