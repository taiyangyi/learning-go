package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// s string 要转换的字符串表示的浮点数
// bitSize int 指定浮点数的位大小。通常使用 32 或 64 来表示单精度和双精度浮点数
// float64 转换后的浮点数值。
// error 如果转换失败，则返回一个非空的错误
// func ParseFloat(s string, bitSize int) (float64, error)

func main() {

	str := "3.1415"

	// 字符串转浮点数(ParseFloat)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换失败：", err)
	}

	fmt.Printf("转换前数据类型：Type of str: %s\n", reflect.TypeOf(str))
	fmt.Println("转换后数据：", num)
	fmt.Printf("转换后数据类型：Type of num: %s\n", reflect.TypeOf(num))
}

// 转换前数据类型：Type of str: string
// 转换后数据： 3.1415
// 转换后数据类型：Type of num: float64
