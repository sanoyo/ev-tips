package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// 2回目以降の出力を配列に入れる
	values := []int{}
	count := 0
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		// 入力された数に、奇数がないか確認
		if number%2 != 0 {
			// 最初に奇数が存在している場合は、0を返す
			if count == 0 {
				fmt.Println(0)
				continue
			}

		} else {
			// 偶数の場合
			values = append(values, number)
			count++
			if count == n {
				fmt.Println(count - 1)
			}
		}
	}
}
