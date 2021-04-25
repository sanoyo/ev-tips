// ref: https://qiita.com/tenntenn/items/686a75e11e8dcd9912ec
// cap は容量、lenは値の数

package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	// 10
	fmt.Println(cap(ch1))

	ch2 := make(chan int)
	// 0
	fmt.Println(cap(ch2))

	ch := make(chan struct{}, 10)
	ch <- struct{}{}
	// cap: 10 len: 1
	fmt.Println("cap:", cap(ch), "len:", len(ch))
	// <-ch
	// cap: 10 len: 0
	fmt.Println("cap:", cap(ch), "len:", len(ch))
}
