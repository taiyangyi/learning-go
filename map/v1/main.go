package main

import "fmt"

// map 键值对的无序集合
// var 变量名 map[键数据类型]值数据类型
// 声明并初始化 var 变量名 = make(map[键数据类型]值数据类型,长度) 长度参数可省略

func main() {

	// 获取值/改变值
	var m = make(map[int]string)
	fmt.Printf("Map 长度 = %d\n", len(m))
	// Map 长度 = 0

	// 为 map 赋值
	m[0] = "China"
	m[1] = "American"
	m[2] = "Russia"

	fmt.Printf("Map 长度 = %d\n", len(m))
	// Map 长度 = 3
	fmt.Println(m)
	// map[0:China 1:American 2:Russia]

	// 删除元素，调用 delete() 方法完成
	// func delete(m map[Type]Type1, key Type)
	// 删除键为1的元素
	delete(m, 1)
	fmt.Printf("Map 长度=%d,m=%v\n", len(m), m)
	// Map 长度=2,m=map[0:China 2:Russia]

	// 判断元素是否存在
	if _, ok := m[1]; !ok {
		fmt.Println("m[1]元素不存在")
	}
	// m[1]元素不存在
}
