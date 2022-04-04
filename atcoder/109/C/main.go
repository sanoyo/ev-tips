package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())
	return t
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	var N, X int
	fmt.Scan(&N, &X)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	A := make([]int, N)
	for i := range A {
		A[i] = nextInt(sc)
	}

	A = append(A, X)
	sort.Ints(A)
	fmt.Println(A)

	intervals := make([]int, 0)
	// 0番目をスキップするように実装
	for i := 1; i < len(A); i++ {
		fmt.Println(i)
		intervals = append(intervals, A[i]-A[i-1])
	}

	fmt.Println(intervals)

	ans := intervals[0]
	fmt.Println(ans)
	for i := 0; i < len(intervals); i++ {
		ans = gcd(ans, intervals[i])
	}
	fmt.Println(ans)
}
