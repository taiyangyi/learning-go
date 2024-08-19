package main

import (
	"fmt"
	"strings"
)

func main() {

	var builder strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&builder, "%d...", i)
	}

	builder.WriteString("hello")
	fmt.Println(builder.String())

}
