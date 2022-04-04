package main

import (
	"fmt"
	"sort"
)

type ivec []int

// コレクション内の要素数
func (v ivec) Len() int { return len(v) }

// インデックスiで示される要素がインデックスjの前に並び替えるべきかを返す
func (v ivec) Less(i, j int) bool { return v[i] > v[j] }

// インデックスi,jの要素を入れ替える
func (v ivec) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

func main() {
	var n int
	fmt.Scan(&n)

	mochi := make(ivec, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mochi[i])
	}

	sort.Sort(mochi)

	count := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			count++
		}
		if i != n-1 && mochi[i] != mochi[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
