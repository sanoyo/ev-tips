// Ref: https://atcoder.jp/contests/arc037/submissions/6252689
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func Search(n int, cond func(i int) bool) int {
	for m := -1; (n - m) > 1; {
		i := int(m + (n-m)/2)
		if cond(i) {
			n = i
		} else {
			m = i
		}
	}
	return n
}

func Solve(in io.Reader, out io.Writer) {
	// n 列数
	// k 求めたい位置
	var n, k int
	var as, bs []int

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ = strconv.Atoi(scanner.Text())

	as = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		as[i], _ = strconv.Atoi(scanner.Text())
	}
	bs = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		bs[i], _ = strconv.Atoi(scanner.Text())
	}

	// 昇順にsort
	sort.Ints(as)
	sort.Ints(bs)

	// 二分探索
	ret := Search(as[n-1]*bs[n-1], func(x int) bool {
		cnt := 0
		for _, a := range as {
			// booleanを返す
			// a*bs[i] > x がスライドで解説指定した箇所
			cnt += Search(n, func(i int) bool { return a*bs[i] > x })
		}
		// booleanを返す
		return cnt >= k
	})
	fmt.Fprintln(out, ret)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
