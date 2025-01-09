package main

import "fmt"

// 切片追加
// 总结：
// 1、切片的长度和容量不同，长度表示左指针至右指针之间的距离，容量表示左指针至底层数组末尾的距离
// 2、切片的扩容机制，append 的时候，如果长度增加后超过容量，则将容量增加2倍

func main() {

	// make(类型、长度、容量)
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// len=3,cap=5,slice=[0 0 0]

	// 向 numbers 切片追加一个元素1
	numbers = append(numbers, 1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// len=4,cap=5,slice=[0 0 0 1]

	// 向 numbers 切片追加一个元素2
	numbers = append(numbers, 2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// len=5,cap=5,slice=[0 0 0 1 2]

	// 向numbers再次追加元素，追加前cap已经达到容量，因为我们是动态数组，他们自己扩容，扩出numbers之前相同的容量
	// 向 numbers 切片追加一个元素3
	numbers = append(numbers, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	//len=6,cap=10,slice=[0 0 0 1 2 3]

}
