package main

import "fmt"

func main() {
	var n, money int
	fmt.Scan(&n, &money)

	for i := 0; i <= n; i++ {
		for j := 0; j <= (n-i); j++ {
			k := n-i-j
			if 10000*i+5000*j+1000*k == money {
				fmt.Println(i, j, k)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
