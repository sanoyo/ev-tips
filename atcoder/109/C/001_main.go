package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	// "sort"
	"strconv"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())
	return t
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

	A = append(A, X) //Xを追加
	sort.Ints(A)
	intervals := make([]int, 0)
	for i := 1; i < len(A); i++ {
		intervals = append(intervals, A[i]-A[i-1])
	}
	fmt.Println(intervals)
	fmt.Println("-----")

	// intervalsの最初の値
	ans := intervals[0]
	for i := 1; i < len(intervals); i++ {
		ans = gcd(ans, intervals[i])
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
