package main

import (
	"fmt"
)

// 上下左右に動くための変数準備
var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

// キューを定義
var q = make([]int, 0)

func bfs(x, y int) {
	for i := range dx {
		nx := dx[i]
		ny := dy[i]
		// 範囲外かどうか かつ . であることを確認
		if 0 <= nx && nx < w && 0 <= ny && ny < h && seamap1[ny][nx] == '.' {
			q[i] += nx
			bfs(ny, nx)
		}
	}
}

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	var sy, sx, gy, gx int
	fmt.Scan(&sy, &sx, &gy, &gx)

	// スタートの座標
	var s = [2]int{sy, sx}
	// ゴールの座標
	var g = [2]int{gy, gx}
	fmt.Println(s)
	fmt.Println(g)

	lines := make([][]byte, r)
	for i := 0; i < r; i++ {
		fmt.Scan(&lines[i])
	}

	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			if lines[y][x] == "." {
				bfs(y, x)
			}
		}
	}
}
