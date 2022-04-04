package main
 
import (
	"fmt"
)

// 834 を 10 で割った余りは 4　-> 答えに加算
// 834 を 10 で割って 83
// 83 を 10 で割った余りは 3 -> 答えに加算
// 83 を 10 で割って 8
// 8 を 10 で割った余りは 8 -> 答えに加算
// 8 を 10 で割って 0
// 0 なので break
 
func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
 
func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
 
	total := 0
	for i := 0; i <= n; i++ {
		if a <= findSumOfDigits(i) && findSumOfDigits(i) <= b {
			total += i
		}
	}
	fmt.Println(total)
}