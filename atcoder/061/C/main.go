package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	var ans int
	length := len(s) - 1
	for bit := 0; bit < 1<<uint(length); bit++ {
		now := string(s[0])
		for i := 0; i < length; i++ {
			if (bit>>i)&1 == 1 {
				now += "+"
			}
			now += string(s[i+1])
		}
		for _, val := range strings.Split(now, "+") {
			t, _ := strconv.Atoi(val)
			ans += t
		}
	}
	fmt.Println(ans)
}
