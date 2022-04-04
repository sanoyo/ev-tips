package main

import (
	"fmt"
)

func cal(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * cal(n-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(cal(n))
}
