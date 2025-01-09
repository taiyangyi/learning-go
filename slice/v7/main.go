package main

import "fmt"

func main() {

	sli := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(sli); i++ {
		fmt.Printf("索引下标：%d, value=%d\n", i, sli[i])
	}

	// 索引下标：0, value=1
	// 索引下标：1, value=2
	// 索引下标：2, value=3
	// 索引下标：3, value=4
	// 索引下标：4, value=5
	// 索引下标：5, value=6

	sli2 := make([]string, 3)
	sli2[0] = "China"
	sli2[1] = "American"
	sli2[2] = "Russia"
	for k, v := range sli2 {
		fmt.Printf("索引下标：%d, value=%v\n", k, v)
	}

	// 	索引下标：0, value=China
	// 索引下标：1, value=American
	// 索引下标：2, value=Russia

}
