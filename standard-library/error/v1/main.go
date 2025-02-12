package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err error = fmt.Errorf("%s", "the error is normal")
	// 等价于下边
	err := fmt.Errorf("%s", "the error is normal")
	fmt.Println("err = ", err)
	// 输出结果：err =  the error is normal

	// 或者直接使用 error 包
	err2 := errors.New("the error is normal")
	fmt.Println("err2 = ", err2)
	// 输出结果：err2 =  the error is normal

}
