// https://atcoder.jp/contests/abc153/submissions/9876070
package main

import (
	"fmt"
	"math"
)

func main() {
	var H int
	fmt.Scan(&H)

	var c int
	for 0 < H {
		H /= 2
		c++
	}
	// 塁乗
	r := int(math.Pow(2, float64(c)))
	fmt.Println(r - 1)
}
