package main

import (
	"fmt"
)

// 1回目
// func nextInt() int {
// 	sc.Scan()
// 	i, e := strconv.Atoi(sc.Text())
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// var sc = bufio.NewScanner(os.Stdin)
// func main() {
// 	sc.Split(bufio.ScanWords)
// 	n := nextInt()
// 	x := 0
// 	for i := 0; i < n; i++ {
// 		x += 1
// 		if nextInt() == 1 {
// 			fmt.Println(x)
// 		}
// 	}
// 	fmt.Println(n)
// }

func main() {
	var n int
	fmt.Scan(&n)
	var c = 0
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if tmp == c+1 {
			c++
		}
	}
	if c == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(n - c)
}
