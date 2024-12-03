package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// str string：要解析的字符串。
// bool, error：返回两个值，一个是解析后的布尔值，另一个是可能的错误。
// 如果字符串解析成功，函数返回转换后的布尔值和 nil 的错误；如果解析失败，
// 函数返回一个非 nil 的错误，其中错误信息描述了解析失败的原因。

// 字符串转布尔值
// func ParseBool(str string) (bool, error)

func main() {

	str := "false"
	bool, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("转换失败！")
	}
	fmt.Printf("转换前数据类型：Type of str: %s\n", reflect.TypeOf(str))
	fmt.Println("转换后数据：", bool)
	fmt.Printf("转换后数据类型：Type of num: %s\n", reflect.TypeOf(bool))
}

// 转换前数据类型：Type of str: string
// 转换后数据： false
// 转换后数据类型：Type of num: bool
