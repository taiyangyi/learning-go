package main

import "fmt"

// 自定义错误结构体
type MyError struct {
	when string
	what string
}

// 实现 error 接口的 Error 方法
func (myError *MyError) Error() string {
	return fmt.Sprintf("%s-%s", myError.when, myError.what)
}

func run() error {

	return &MyError{
		when: "just now",
		what: "something went wrong",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
