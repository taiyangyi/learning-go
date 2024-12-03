package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// strconv.Atoi 用于将字符串转换为整数
// func Atoi(s string) (int, error)

func main() {

	str := "999"

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	fmt.Printf("转换前数据类型：Type of str: %s\n", reflect.TypeOf(str))
	fmt.Println("转换后数据：", num)
	fmt.Printf("转换后数据类型：Type of num: %s\n", reflect.TypeOf(num))

}

// 输出结果

// 转换前数据类型：Type of str: string
// 转换后数据： 999
// 转换前数据类型：Type of num: int
