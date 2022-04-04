package main

import (
	"fmt"
)

func main() {
	var s, t, u string
	var a, b int
	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&u)

	if a == 0 && b == 0 {
		fmt.Println(a, b)
		return
	}

	if a > 0 && u == s {
		a = a - 1
	} else if b > 0 && u == t {
		b = b - 1
	}

	fmt.Println(a, b)
}
