package main

import "fmt"

func main() {

	m := make(map[int]string)
	m[0] = "振铃"
	m[1] = "摘机"
	m[2] = "空闲"
	m[3] = "异常"

	// 遍历map，首先map遍历是无序的，不论是键还是值，
	// 如果需要遍历时保持相同的顺序，提要将键做排序处理
	for k, v := range m {
		fmt.Printf("key = %d,val = %s\n", k, v)
	}

	// 多次运行，每次输出的结果顺序不一致

	// key = 0,val = 振铃
	// key = 1,val = 摘机
	// key = 2,val = 空闲
	// key = 3,val = 异常

	// key = 3,val = 异常
	// key = 0,val = 振铃
	// key = 1,val = 摘机
	// key = 2,val = 空闲

}
