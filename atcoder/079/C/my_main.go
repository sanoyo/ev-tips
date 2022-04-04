package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cal(sum int, s string) int {
	count := sum
	plus := 0
	all := ""
	total := ""
	for i := 1; i <= len(s); i++ {
		n, _ := strconv.Atoi(s[plus:i])
		count = n + count
		plus++
		all = strconv.Itoa(n) + "+"
		total = total + all
		if len(s) == i {
			total := strings.Replace(total, "+", "=", 2)
			fmt.Printf(total + strconv.Itoa(count))
			// fmt.Printf("=%d", count)
		}
	}

	return count
}

func main() {
	var s string
	fmt.Scan(&s)
	cal(0, s)
}
