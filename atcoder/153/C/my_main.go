package main

import "fmt"

func remove(numbers []int, search int) []int {
	result := []int{}
	for _, num := range numbers {
		if num != search {
			result = append(result, num)
		}
	}
	return result
}

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
	var n, k int
	fmt.Scan(&n, &k)

	tec := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tec[i])
	}

	remove(tec, 1)
	fmt.Println(tec)

	// kの数だけ、最大値順に配列を減らす
	// for i := 0; i < k; i++ {
	// 	remove(tec, max(tec))
	// 	fmt.Println(tec)
	// }

	// 残りの配列を1ずつ減らしていく
	// count := 0
	// for i := 0; i < len(tec); i++ {
	// 	n--
	// 	count++
	// }
	// fmt.Println(count)
}
