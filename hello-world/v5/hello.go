package main

import "fmt"

const helloPrefix = "Hello,"

// 当我们的函数用空字符串调用时，它默认为打印 「Hello, World」
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("")) // Hello,World
}
