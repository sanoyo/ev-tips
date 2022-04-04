package main

import (
	"fmt"
)

var move1 = 1
var move2 = 2

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	coordinate := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&coordinate[i])
	}

	// var diff float64
	var a int
	var next_a int
	for i := 0; i < n; i++ {
		// diff = math.Abs(float64(x - coordinate[i]))
		a = x % coordinate[i]
		if next_a != 0 && i != n-1 {
			next_a = x % coordinate[i+1]
			if a == next_a {
				break
			}
		}
	}
}
