package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, K int
	fmt.Scan(&n, &K)

	// 2回目以降以降の出力を配列に入れる
	values := []int{}
	for i := 0; i < n*2; i++ {
		var number int
		fmt.Scan(&number)
		values = append(values, number)
	}

	// 昇順にする
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	// Kの値を見つけ出す
	fmt.Println(values[K-1])
}
