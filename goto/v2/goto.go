package main

import "fmt"

func main() {
	i := 0
	i++
	if i < 5 {
		goto Here // 跳转到标签Here
	}
Here: // 标签
	fmt.Println(i)
}
