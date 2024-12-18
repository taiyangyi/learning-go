package main

import (
	"fmt"
	"log"

	"github.com/taiyangyi/learning-go/datafile"
)

// 投票统计

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	fmt.Println(counts)

}
