package main

import "fmt"

// slice 切片数组

// 声明
// var sli []int    var 变量名 []数据类型

// 声明及初始化
// var 变量名 = make([]int,长度,容量) 其中，容量参数可以省略
// var sli = make([]int,5,10)

func main() {

	// 声明一个整型切片，但是不分配空间
	var sli []int
	fmt.Println(sli)
	fmt.Printf("切片长度=%d,切片容量=%d\n", len(sli), cap(sli))
	// 	[]
	// 切片长度=0,切片容量=0

	// 声明及初始化
	sli = make([]int, 3)
	fmt.Println(sli)
	fmt.Printf("切片长度=%d,切片容量=%d\n", len(sli), cap(sli))
	// 	[0 0 0]
	// 切片长度=3,切片容量=3

	// 为 sli 切片元素赋值
	sli[0] = 1
	sli[1] = 2
	sli[2] = 3
	fmt.Println(sli)
	// [1 2 3]
}
