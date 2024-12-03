package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// b bool：要转换的布尔值。
// string：对应布尔值的字符串表示。可能是 "true" 或 "false"
// func FormatBool(b bool) string

func main() {

	boolstr := true
	// 布尔值转字符串
	str := strconv.FormatBool(boolstr)
	fmt.Printf("转换前数据类型：Type of str: %s\n", reflect.TypeOf(boolstr))
	fmt.Println("转换后数据：", str)
	fmt.Printf("转换后数据类型：Type of num: %s\n", reflect.TypeOf(str))
}

// 转换前数据类型：Type of str: bool
// 转换后数据： true
// 转换后数据类型：Type of num: string
