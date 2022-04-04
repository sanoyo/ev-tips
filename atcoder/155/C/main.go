package main

import (
	"fmt"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)

	texts := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&texts[i])
	}

	var num int
	ans := make([]string, 0)
	maps := make(map[string]int, n)
	// mapで同じものをカウント
	for i := 0; i < n; i++ {
		maps[texts[i]]++
		num = max(num, maps[texts[i]])
	}

	// value(数値)とmaxのカウント数が同じなのであれば、ansに追加
	for key, value := range maps {
		if value == num {
			ans = append(ans, key)
		}
	}
	sort.Strings(ans)

	for _, text := range ans {
		fmt.Println(text)
	}
}
