package main

import "fmt"

func main() {
	// 定义一个标签 Start
Start:
	fmt.Println("Start of the program")
	// 使用 goto 语句跳转到标签 end
	goto End
	fmt.Println("unreachable")

End:
	fmt.Println("End of the program")
	goto Start

}
