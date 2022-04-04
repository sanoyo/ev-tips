package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scan(&n)

	// fmt.Println(n[:len(n)])   //123
	// fmt.Println(n[:len(n)-1]) //12
	// fmt.Println(n[1:3])       //23
	// fmt.Println(n[:len(n)-2]) //1
	// fmt.Println(n[1:2])       //2
	// fmt.Println(n[2:3])       //3
	// fmt.Println("-----")

	count := 0
	count2 := 0
	count3 := 0
	count4 := 0
	for i := 1; i <= len(n); i++ {
		count2, _ = strconv.Atoi(n[:i])
		fmt.Println(count3)

		count3, _ = strconv.Atoi(n[i:])
		fmt.Println(count2)

		// 1+2+5
		count4, _ = strconv.Atoi(n[i-1 : i])

		if count3 != count2 || count2 != 0 || count3 != 0 {
			count += count2 + count3
			count += count4
		}
	}
	fmt.Println(count)
}
