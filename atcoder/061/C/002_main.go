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
	for i := 0; i < 1<<uint(length); i++ {
		now := string(s[0])
		for j := 0; j < length; j++ {
			if i>>uint(j)&1 == 1 {
				now += "+"
			}
			now += string(s[j+1])
		}

		for _, val := range strings.Split(now, "+") {
			t, _ := strconv.Atoi(val)
			ans += t
		}
	}
	fmt.Println(ans)
}
