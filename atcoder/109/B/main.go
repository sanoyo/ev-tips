package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	// mapで重複しているものをチェック
	w := make([]string, n)
	table := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scanln(&w[i])
		table[w[i]] += 1
		// 重複するとtable[w[i]]の値が、2以上になるので、それをチェック
		if table[w[i]] > 1 {
			fmt.Println("No")
			return
		}
	}
	// len(w[i]) - 1　→　プログラミング上の文字の長さ(0から数える)
	for i := 0; i < n-1; i++ {
		fmt.Println(w)
		fmt.Println(w[i])
		fmt.Println(w[i][0])
		fmt.Println(w[i][len(w[i])-1])
		fmt.Println("----------")
		fmt.Println(w[i+1][0])
		if w[i][len(w[i])-1] != w[i+1][0] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
