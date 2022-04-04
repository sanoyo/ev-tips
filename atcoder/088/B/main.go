package main

import (
	"fmt"
	"sort"
)

type ivec []int

func (v ivec) Len() int           { return len(v) }
func (v ivec) Less(i, j int) bool { return v[i] > v[j] }
func (v ivec) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

func main() {
	var n int
	fmt.Scan(&n)
	a := make(ivec, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Sort(a)

	alice := 0
	bob := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}

	fmt.Println(alice - bob)
}
