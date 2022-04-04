// https://atcoder.jp/contests/abc007/submissions/9642287
package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	x, y int
}

// ↑→↓←
var dx [4]int = [4]int{0, 1, 0, -1}
var dy [4]int = [4]int{-1, 0, 1, 0}

func main() {
	var r, c, sy, sx, gy, gx int
	fmt.Scan(&r, &c, &sy, &sx, &gy, &gx)

	// ステージ設定
	stage := make([][]string, r)
	for i := 0; i < r; i++ {
		var tmp string
		fmt.Scan(&tmp)
		stage[i] = strings.Split(tmp, "")
	}
	// fmt.Println(stage)
	// [[# # # # # # # #] [# . . . . . . #] [# . # # # # # #] [# . . # . . . #] [# . . # # . . #] [# # . . . . . #] [# # # # # # # #]]

	// 1 引いているのは、プログラミング上での位置を正確に示すため
	// (2,2)→(1,1)
	stage[sy-1][sx-1] = "S"
	stage[gy-1][gx-1] = "G"

	// キュー
	que := make([]Pair, 0)
	que = append(que, Pair{x: sx - 1, y: sy - 1})

	// ステップ数
	dist := make([][]int, r)
	for i := 0; i < r; i++ {
		dist[i] = make([]int, c)
		for j := 0; j < c; j++ {
			dist[i][j] = -1
		}
	}
	// fmt.Println(dist)
	// [[-1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1]]

	// スタート地点は0にする
	dist[sy-1][sx-1] = 0

	// solve
	for len(que) > 0 {
		point := que[0]
		que = que[1:]

		if point.x == gx-1 && point.y == gy-1 {
			break
		}

		// 上下左右に移動
		for i := 0; i < 4; i++ {
			nx, ny := point.x+dx[i], point.y+dy[i]
			// dist[ny][nx] != -1 →　「探索済みなら」という意味
			if stage[ny][nx] == "#" || dist[ny][nx] != -1 {
				continue
			}

			if 0 <= nx && nx < c && 0 <= ny && ny < r {
				que = append(que, Pair{x: nx, y: ny})
				dist[ny][nx] = dist[point.y][point.x] + 1
			}
		}
	}

	fmt.Println(dist[gy-1][gx-1])
}
