package main

import "fmt"

// 切片截取

func main() {

	num := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始数据：%v,len=%d,cap=%d\n", num, len(num), cap(num))
	// 原始数据：[1 2 3 4 5],len=5,cap=5

	// 截取 范围 左闭右开
	extract := num[0:2]
	fmt.Printf("截取后数据：%v,len=%d,cap=%d\n", extract, len(extract), cap(extract))
	// 截取后数据：[1 2],len=2,cap=5

}
