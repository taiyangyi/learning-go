// 整数转换为字符串
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 整数转字符串
// func Itoa(i int) string

func main() {
	num := 9
	atferConv := strconv.Itoa(num)
	fmt.Printf("转换前数据类型：Type of str: %s\n", reflect.TypeOf(num))
	fmt.Println("转换后数据：", atferConv)
	fmt.Printf("转换后数据类型：Type of num: %s\n", reflect.TypeOf(atferConv))

}

// 转换前数据类型：Type of str: int
// 转换后数据： 9
// 转换后数据类型：Type of num: string
