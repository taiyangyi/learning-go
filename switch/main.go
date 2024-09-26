package main

import "fmt"

func main() {
	fruit := "apple"
	switchofgeneral(fruit)

	switchofexpression()

	switchofdefault()

	switchofomitexpression()

	switchofmorecase()

	switchoffallthrough()

	switchoftypeassert()
}

// 普通表达式
func switchofgeneral(fruit string) {
	switch fruit {
	case "apple":
		fmt.Printf("fruit: %s\n", fruit)
	case "banana":
		fmt.Printf("fruit: %s\n", fruit)
	case "orange":
		fmt.Println("orange")
	}
}

// go run main.go
// 输出：
// fruit: apple

// 运算表达式
func switchofexpression() {
	n := 4
	switch n * 2 {
	case 8:
		fmt.Println("8")
	case 10:
		fmt.Println("10")
	}
}

// go run main.go
// 输出：
// 8

// default
func switchofdefault() {
	n := 1024
	switch n * 2 {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("1024")
	}
}

// go run main.go
// 输出：
// 1024

// 省略表达式
func switchofomitexpression() {
	n := 1024
	switch {
	case n < 1024:
		fmt.Println("n < 1024")
	case n > 1024:
		fmt.Println("n > 1024")
	case n == 1024:
		fmt.Println("n==1024")
	default:
		fmt.Println("无结果")
	}

}

// go run main.go
// 输出：
// n==1024

// 同时多个case
func switchofmorecase() {
	n := 1024
	switch n {
	case 1023, 1024:
		fmt.Println("n <= 1024")
	case 1025:
		fmt.Println("n > 1024")
	default:
		fmt.Println("无结果")
	}
}

// go run main.go
// 输出：
// n <= 1024

func switchoffallthrough() {
	n := 1024
	switch {
	case n < 1024:
		fmt.Println("n < 1024")
		fallthrough // 继续向下执行
	case n > 1024:
		fmt.Println("n > 1024")
		fallthrough // 继续向下执行
	case n == 1024:
		fmt.Println("n == 1024")
		fallthrough // 继续向下执行
	default:
		fmt.Println("无结果")
	}
}

// n == 1024
// 无结果

// 类型断言
func switchoftypeassert() {
	var n any = 1024
	switch n.(type) {

	case int:
		println("n is a int")
	case float64:
		println("n is a float64")
	case bool:
		println("n is a bool")
	case string:
		println("n is a string")
	default:
		println("n is invalid")
	}
}

// n is a int
