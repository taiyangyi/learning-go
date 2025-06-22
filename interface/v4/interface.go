package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Hero 声明 Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// HeroSlice 声明一个 Hero 结构体切片类型
type HeroSlice []Hero

// Len 实现 Interface 接口，要对 Hero 进行排序
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定你使用什么标准进行排序
// 按 Hero 的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

// 交换
func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {
	// 先定一个数组或者切片
	var intSlice = []int{0, -1, 10, 7, 2}
	// 要求对 intSlice 切片进行排序
	// 1 冒泡排序
	// 2 使用 Go 提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 对结构体切片进行排序
	// 1 冒泡排序
	// 2 使用 Go 提供的方法

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	// 排序前的排序
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println("================")
	sort.Sort(heroes)

	// 排序后的排序
	for _, v := range heroes {
		fmt.Println(v)
	}

}
