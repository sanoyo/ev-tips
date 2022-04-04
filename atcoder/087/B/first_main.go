package main

import (
	"fmt"
)

func check(total int, n int, flag bool) int {
	if total == n {
		flag = false
		return total
	}

	return total
}

// 最初に考えたアルゴリズム
func main() {
	var (
		n     int
		total int
		count int
	)

	nums := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&nums[i])
	}

	// 目的の合計金額
	fmt.Scan(&n)

	five_coin := nums[0]
	one_coin := nums[1]
	fifty_coin := nums[2]

	flag := true

	for i := 0; i < five_coin; i++ {
		total += five_coin
		check(total, n, flag)
		count++
		for i := 0; i < one_coin; i++ {
			total += one_coin
			count++
			for i := 0; i < fifty_coin; i++ {
				total += fifty_coin
				count++
			}
		}
	}

	fmt.Println(count)
}
