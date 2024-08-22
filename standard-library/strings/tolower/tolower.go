package main

import (
	"fmt"
	"strings"
)

// strings.tpoLower(s string) string
// ToLower(s string) string：将一个字符串里的大写字符转成小写，因为字符串不可变的特点，该函数会返回一个新的字符串。
func main() {
	s := "BIG HOUSE"
	s_tolower := strings.ToLower(s)
	fmt.Println(s_tolower) // big house
}
