package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	var result string
	for i := 0; i < len(s); i++ {
		result += strings.Replace(s, s, "x", -1)
	}
	fmt.Println(result)
}
