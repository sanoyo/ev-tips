package main

import (
	"fmt"
	"strconv"
)

func dfs(sum int, s string) int {
	res := sum
	for idx := 1; idx < len(s); idx++ {
		n, _ := strconv.Atoi(s[0:idx])
		res += dfs(sum+n, s[idx:])
	}
	n, _ := strconv.Atoi(s)
	return res + n
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(dfs(0, s))
}
