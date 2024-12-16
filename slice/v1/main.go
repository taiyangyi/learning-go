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

	// 声明及初始化，通过make开辟空间容量
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

	// 基于现有数组得到一个切片
	array := [5]int{1, 2, 3, 4, 5}
	sli2 := array[2:4]
	fmt.Printf("数组 array 中 array:%v,len:%v,cap:%v\n", array, len(array), cap(array))
	fmt.Printf("切片 sli2 中 sli2:%v,len:%v,cap:%v\n", sli2, len(sli2), cap(sli2))
	// 	数组 array 中 array:[1 2 3 4 5],len:5,cap:5
	// 切片 sli2 中 sli2:[3 4],len:2,cap:3
}
