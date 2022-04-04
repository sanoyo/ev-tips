package main
 
import (
	"fmt"
	"math"
)
 
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
 
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
 
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	dmin := math.MaxInt32
	ans := 0
	for bit := 0; bit < (1 << uint(n)); bit++ { // 肉の数でbit全探索
		var A, B int
		for j := 0; j < n; j++ {
			if bit&(1<<uint(j)) != 0 { // bitが立って入れば焼き台Aで焼く
				A += a[j]
			} else {
				B += a[j]
			}
		}
		d := abs(A - B)
		if dmin > d {
			dmin = d
			ans = max(A, B) // 差分が最も小さい時の大きい方が答え
		}
	}
	fmt.Println(ans)
}