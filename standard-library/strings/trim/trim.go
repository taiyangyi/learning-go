package main

import (
	"fmt"
	"strings"
)

func main() {

	//---------------------TrimSpace----------------------------//

	// func TrimSpace(s string) string
	// 去掉字符串左右两边的空格
	str := "  hot coffee or hot tea  "
	trimSpaceStr := strings.TrimSpace(str)
	fmt.Println("改造前:", str)
	fmt.Println("改造后:", trimSpaceStr)

	// 改造前:   hot coffee or hot tea
	// 改造后: hot coffee or hot tea

	//---------------------Trim----------------------------//

	// func Trim(s, cutset string) string
	// 将字符串[左右两边]所指定的字符（串）去掉

	// 第一个参数 s 为需要去除指定字符的字符串。
	// 第二个参数为指定的字符（串）。

	str2 := "=I like my blue house="
	str2Trim := strings.Trim(str2, "=")
	fmt.Println("改造前:", str2)
	fmt.Println("改造后:", str2Trim)

	// 改造前: =I like my blue house=
	// 改造后: I like my blue house

	str2_1 := "=I like my blue house=!#$"
	str2_1Trim := strings.Trim(str2_1, "=!#$")
	fmt.Println("改造前:", str2_1)
	fmt.Println("改造后:", str2_1Trim)

	// 	改造前: =I like my blue house=!#$
	// 改造后: I like my blue house

	//-----------------------TrimRight--------------------------//

	// func TrimRight(s, cutset string) string
	// 	将字符串[右边]所指定的字符（串）去掉
	// 	第一个参数 s 为需要去除指定字符的字符串。
	// 第二个参数为指定的字符（串）。

	str3 := "=I like my blue house="
	str3TrimRight := strings.TrimRight(str3, "=")
	fmt.Println("改造前:", str3)
	fmt.Println("改造后:", str3TrimRight)

	// 	改造前: =I like my blue house=
	// 改造后: =I like my blue house

	//-----------------------TrimLeft--------------------------//

	// func TrimLeft(s, cutset string) string
	// 	将字符串[左边]所指定的字符（串）去掉
	// 第一个参数 s 为需要去除指定字符的字符串。
	// 第二个参数为指定的字符（串）。

	str4 := "=I like my blue house="
	str4TrimLeft := strings.TrimLeft(str4, "=")
	fmt.Println("改造前:", str4)
	fmt.Println("改造后:", str4TrimLeft)

	// 	改造前: =I like my blue house=
	// 改造后: I like my blue house=

	//-----------------------TrimPrefix--------------------------//
	// func TrimPrefix(s, prefix string) string
	// TrimPrefix 返回没有提供前缀字符串的 s。如果 s 不以前缀开头，则不变地返回 s。
	str5 := "====@@@I like my blue house="
	str5TrimPrefix := strings.TrimPrefix(str5, "====@@@")
	fmt.Println("改造前:", str5)
	fmt.Println("改造后:", str5TrimPrefix)

	// 	改造前: ====@@@I like my blue house=
	// 改造后: I like my blue house=

	str5_2_TrimPrefix := strings.TrimPrefix(str5, "====@@I like ")
	fmt.Println("改造后:", str5_2_TrimPrefix)
	// 改造后: ====@@@I like my blue house=
	// 如果没有对应性的前缀字符串，则返回原字符串。

	//-----------------------TrimPrefix--------------------------//

	// func TrimSuffix(s, suffix string) string
	// 返回没有提供后缀字符串的 s。如果 s 没有以后缀结束，则不变地返回 s。
	str6 := "====@@@I like my blue house="
	str6TrimSuffix := strings.TrimSuffix(str6, "=")
	fmt.Println("改造前:", str6)
	fmt.Println("改造后:", str6TrimSuffix)

	// 	改造前: ====@@@I like my blue house=
	// 改造后: ====@@@I like my blue house

}
