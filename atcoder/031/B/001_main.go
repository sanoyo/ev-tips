// https://atcoder.jp/contests/arc031/submissions/5810403
package main

import "fmt"

var (
	C      [10][10]rune
	N      int // the number of lands(o)
	marked int
	dy     = []int{1, 0, -1, 0}
	dx     = []int{0, 1, 0, -1}
)

func reset() {
	marked = 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if C[i][j] == '#' {
				C[i][j] = 'o'
			}
		}
	}
}

func dfs(y, x int) bool {
	if marked == N {
		return true
	}
	C[y][x] = '#'
	// テンプレートでいうtrueの役割
	marked++
	for i := 0; i < len(dy); i++ {
		ny := y + dy[i]
		nx := x + dx[i]
		if 0 <= nx && nx < 10 && 0 <= ny && ny < 10 && C[ny][nx] == 'o' {
			if dfs(ny, nx) {
				return true
			}
		}
	}
	return false
}

func main() {
	// ○の数をカウントする
	// C →　10行
	// C[i] →　10行のうちの1行
	// C[i][j] →　1行のうち1文字
	for i := 0; i < 10; i++ {
		var s string
		fmt.Scan(&s)
		for j := 0; j < 10; j++ {
			C[i][j] = rune(s[j])
			if C[i][j] == 'o' {
				// ○の数
				N++
			}
		}
	}

	// DFSを適応
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// oだった場合は、もう一度for文の最初に戻る
			// このoは、テンプレでいうところのtrue??
			if C[i][j] == 'o' {
				continue
			}
			// 必要ない
			// C[i][j] = 'o'

			// C[i][j] == 'x'の場合、探索を実施する
			// dfsメソッドの返り値が、trueであれば返す
			if dfs(i, j) {
				fmt.Println("YES")
				return
			}
			C[i][j] = 'x'
			reset()
		}
	}
	fmt.Println("NO")
}
