package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)
	// 2024-12-19 10:14:44.9214567 +0800 CST m=+0.000058001
	fmt.Printf("%2d.%2d.%4d\n", t.Day(), t.Month(), t.Year())
	// 19.12.2024

	t = time.Now().UTC()
	fmt.Println(time.Now())
	// 2024-12-19 10:19:49.5286257 +0800 CST m=+0.024544501
	fmt.Println(t)
	// 2024-12-19 02:19:49.5286257 +0000 UTC

}
