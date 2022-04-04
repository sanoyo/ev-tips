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
	var h, w, n int
	// var sy, sx, gy, gx int
	fmt.Scan(&h, &w, &n)
	// fmt.Scan(&sy, &sx, &gy, &gx)

	// ステージ設定
	stage := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Scan(&tmp)
		stage[i] = strings.Split(tmp, "")
		for j := 0; j < h; j++ {
			// スタート地点を-3に置き換え
			if stage[i][j] == "S" {
				stage[i][j] = "0"
			}
		}
	}

	// // キュー
	// que := make([]Pair, 0)
	// que = append(que, Pair{x: sx - 1, y: sy - 1})

	// ステップ数
	// 全て -1 にする
	dist := make([][]string, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]string, w)
		for j := 0; j < w; j++ {
			dist[i][j] = "-1"
		}
	}

	// // スタート地点は0にする
	// dist[sy-1][sx-1] = 0

	// // bfs
	// for len(que) > 0 {
	// 	point := que[0]
	// 	que = que[1:]

	// 	if point.x == gx-1 && point.y == gy-1 {
	// 		break
	// 	}

	// 	// 上下左右に移動
	// 	for i := 0; i < 4; i++ {
	// 		nx, ny := point.x+dx[i], point.y+dy[i]
	// 		// dist[ny][nx] != -1 →　「探索済みなら」という意味
	// 		if stage[ny][nx] == "X" || dist[ny][nx] != -1 {
	// 			continue
	// 		}

	// 		if 0 <= nx && nx < w && 0 <= ny && ny < h {
	// 			que = append(que, Pair{x: nx, y: ny})
	// 			dist[ny][nx] = dist[point.y][point.x] + 1
	// 		}
	// 	}
	// }

	// fmt.Println(dist[gy-1][gx-1])
}
