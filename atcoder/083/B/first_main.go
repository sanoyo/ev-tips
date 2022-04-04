package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	nums := []int{}

	for i := 0; i <= n; i++ {
		nums = append(nums, i)

		if a < nums[i] && nums[i] < b {
			// 1桁
			num := 121 / 10
			dig := 121 % 10
			// fmt.Println(nums[i])
			fmt.Println(dig)
			fmt.Println(num)
		}
	}
}
