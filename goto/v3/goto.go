package main

import "fmt"

func main() {
Outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				goto Outer // 跳出外层循环
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
