package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	contents := make([]string, n)
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&contents[i])
		m[contents[i]]++
	}

	fmt.Println(m)

	// for i := 0; i < n; i++ {
	// 	if contents[i] == contents[i+1] {
	// 		fmt.Println("No")
	// 		return
	// 	}
	// }
	// fmt.Println("Yes")
}
