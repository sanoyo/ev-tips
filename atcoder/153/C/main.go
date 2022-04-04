// https://atcoder.jp/contests/abc153/submissions/9756127
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	hs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&hs[i])
	}

	if n <= k {
		fmt.Print(0)
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(hs)))

	var hs2 []int
	if k == 0 {
		hs2 = hs
	} else {
		hs2 = hs[k:]
	}

	cal := 0
	for _, h := range hs2 {
		cal = cal + h
	}

	fmt.Print(cal)
}
