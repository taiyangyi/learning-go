package main

import (
	"fmt"
	"strings"
)

// strings.tpoUpper(s string) string
// 将一个字符串里的小写字符转成大写，因为字符串不可变的特点，该函数会返回一个新的字符串。
func main() {

	s1 := "hello world"
	s1_toupper := strings.ToUpper(s1)
	fmt.Println(s1_toupper) // HELLO WORLD
}
