package main
 
import (
	"fmt"
)
 
func main() {
	var (
		n     int
		count int
	)
	fmt.Scan(&n)
	// N個分のスライスを定義
	nums := make([]int, n)
	// N回文数値を受けとる
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
 
	flag := true
	// flagがtrueだったら、for分が回り続ける
	for flag {
		for i := 0; i < n; i++ {
			if nums[i]%2 == 0 {
				nums[i] /= 2
			} else {
				flag = false
			}
		}
		if flag {
			count++
		}
	}
	fmt.Println(count)
}
