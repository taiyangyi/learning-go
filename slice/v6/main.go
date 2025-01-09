package main

import "fmt"

func main() {
	sli := make([]int, 3)
	fmt.Printf("%v ptr:%p\n", sli, sli)
	// [0 0 0] ptr:0xc000016120

	sli2 := sli
	fmt.Printf("%v ptr:%p\n", sli, sli)
	// [0 0 0] ptr:0xc000016120

	sli2[1] = 20
	fmt.Printf("%v ptr:%p\n", sli, sli)
	// [0 20 0] ptr:0xc000016120
}
