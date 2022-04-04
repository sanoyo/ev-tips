package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	ws := make([]string, n)
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		var w string
		fmt.Scan(&w)
		ws[i] = w
		m[w]++
	}

	var last string
	for _, v := range ws {
		ss := strings.Split(v, "")
		fmt.Println(ss)
		fmt.Println(m[v])
		if (last != "" && ss[0] != last) || m[v] > 1 {
			fmt.Println("No")
			return
		}
		last = ss[len(ss)-1]
		fmt.Println(last)
	}
	fmt.Println("Yes")
}
