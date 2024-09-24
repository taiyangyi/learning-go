package main

import "fmt"

func main() {
	//demo()
	// demo2()
	// demo3()
	demo4()
}

// key, val 都保留
func demo() {
	str := "helloworld"
	for k, v := range str {
		fmt.Printf("k=%d, v=%c\n", k, v)
	}
}

// k=0, v=h
// k=1, v=e
// k=2, v=l
// k=3, v=l
// k=4, v=o
// k=5, v=w
// k=6, v=o
// k=7, v=r
// k=8, v=l
// k=9, v=d

// 省略 key
func demo2() {
	str := "helloworld"
	for _, v := range str {
		fmt.Printf("v = %c\n", v)
	}
}

// v = h
// v = e
// v = l
// v = l
// v = o
// v = w
// v = o
// v = r
// v = l
// v = d

// 省略 val
func demo3() {
	str := "helloworld"
	for k, _ := range str {
		fmt.Printf("k = %d\n", k)
	}
}

// k = 0
// k = 1
// k = 2
// k = 3
// k = 4
// k = 5
// k = 6
// k = 7
// k = 8
// k = 9

// key, val 都省略
func demo4() {
	str := "helloworld"
	for _, _ = range str {
		println("do something")
	}
}

// do something
// do something
// do something
// do something
// do something
// do something
// do something
// do something
// do something
// do something
