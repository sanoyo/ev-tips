package main

import (
	"fmt"
	"sort"
)

func maxInt(card []int) int {
	sort.Sort(sort.IntSlice(card))
	return card[len(card)-1]
}

func remove(numbers []int, search int) []int {
	result := []int{}
	for _, num := range numbers {
		// 該当する値(search)以外をスライスとして入れ直す
		if num != search {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	card := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&card[i])
	}

	alice := 0
	bob := 0
	for i := 1; i <= n; i++ {
		// fmt.Println(i)
		if i%2 != 0 {
			max := maxInt(card)
			alice += max
			card = remove(card, max)
		} else {
			max := maxInt(card)
			bob += max
			card = remove(card, max)
		}
	}
	fmt.Println(alice - bob)
}
