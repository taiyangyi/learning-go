package main

import "fmt"

func main() {

	// map类型不支持零值可用，未显式赋值的 map 类型变量的零值为 nil
	// 对处于零值状态的 map变量进行操作将会导致运行时 panic
	// var m map[string]int
	// m["key1"] = 1
	//fmt.Println(m) // panic: assignment to entry in nil map

	// map 变量必须显式初始化才能使用，有两种方式：1、使用复合字面值；2、使用 make 预声明
	// 1、使用复合字面值创建 map 类型变量
	var text = map[int]string{
		1: "OK",
		2: "Created",
		3: "Accepted",
	}
	fmt.Println("map的值为:", text) // map的值为: map[1:OK 2:Created 3:Accepted]

	// 2、使用 make进行创建
	m := make(map[string]string)
	m["中国"] = "China"
	m["俄罗斯"] = "Russia"
	m["美国"] = "American"
	fmt.Println("使用 make 创建的 map:", m)
	// 使用 make 创建的 map: map[中国:China 俄罗斯:Russia 美国:American]

}
