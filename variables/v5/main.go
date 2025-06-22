package main

import (
	"fmt"
	"net"
)

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:8080")
	// 如果不想接收err值的话，可以使用 _ 表示，此应用被称为：匿名变量
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(conn)
}
