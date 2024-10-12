package main

import "fmt"

// 数组
// 语法规则：
// 1、var 变量名称 [数组长度]数据类型
// 如： var arr [6]int

// 2、var arr [...]int{v1,v2,v3}
// 在确定数组元素的具体数量以及值时，可以省略 长度 变量，这种方式称为 数组常量。

func main() {

	// 数组的元素可以通过 索引 来获取或修改，索引 从 0 开始。
	var arr [3]string

	arr[0] = "China"
	arr[1] = "American"
	arr[2] = "Britain"

	fmt.Println("中国", arr[0])
	fmt.Printf("美国%s", arr[1])
	fmt.Print("\n", arr[2], "\n")

	// 中国 China
	// 美国American
	// Britain

	//  使用 数组常量 定义
	var numArr = [...]int{1, 2, 3}
	fmt.Println(numArr[0])
	fmt.Println(numArr[1])
	fmt.Println(numArr[2])

	// 1
	// 2
	// 3

	// 获取数组的长度
	fmt.Printf("获取数组长度：%d\n", len(numArr)) // 获取数组长度：3

	// 遍历数组

	// 1、普通循环
	for i := 0; i < len(numArr); i++ {
		fmt.Printf("index = %d, value = %d\n", i, numArr[i])
	}

	// index = 0, value = 1
	// index = 1, value = 2
	// index = 2, value = 3

	// 2、range 循环
	for i, v := range arr {
		fmt.Printf("索引 = %d, 值 = %s\n", i, v)
	}

	// 索引 = 0, 值 = China
	// 索引 = 1, 值 = American
	// 索引 = 2, 值 = Britain

}
