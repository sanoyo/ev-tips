package main

import "fmt"

func max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	var n int
	fmt.Scan(&n)

	var t = make([]int, n)

	if n == 1 {
		fmt.Println(n)
	}

	max := max(t)
	one := t[0] + max
	two := t[1] + t[2]
}
