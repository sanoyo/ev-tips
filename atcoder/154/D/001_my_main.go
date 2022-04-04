package main

import (
	"fmt"
)

func main() {
	var n, k float64
	fmt.Scan(&n, &k)

	a := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var result float64
	x := 0.1667
	for i := 0; i < n; i++ {
		result += a[i] * x
	}
	fmt.Println(result)
}
