package main

import (
	"fmt"
	"strings"
)

// 高效拼接字符串
// 使用 strings 库里的 Builder 变量，结合其写入方法如 WriteString 方法，可以进行高效的拼接字符串。
func main() {

	var builder strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&builder, "%d...", i)
	}

	builder.WriteString("hello")
	fmt.Println(builder.String()) // // 3...2...1...hello

	builder.WriteString("world")
	builder.WriteString("!")
	s := builder.String()
	fmt.Println(s) // 3...2...1...helloworld!

}
