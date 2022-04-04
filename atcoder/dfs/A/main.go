// Ref: https://atcoder.jp/contests/atc001/submissions/6242117
package main

import "fmt"

var h, w int
var sx, sy int       // スタートの座標
var gx, gy int       // ゴールの座標
var reached [][]bool // 座標(x,y)に到達したか
var maze [][]byte    // 迷路の入力を受け取る配列()

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, -1, 1}

// DFSによる深さ優先探索
func dfs(x, y int) {
	if x < 0 || w <= x || y < 0 || h <= y || maze[y][x] == '#' {
		return
	}
	if reached[y][x] {
		return
	}
	reached[y][x] = true
	for i := 0; i < 4; i++ {
		dfs(x+dx[i], y+dy[i])
	}
}

func main() {
	fmt.Scan(&h, &w)

	maze = make([][]byte, h)
	reached = make([][]bool, h)

	for i := range maze {
		maze[i] = make([]byte, w)
		reached[i] = make([]bool, w)
		fmt.Scan(&maze[i])
		for j := range maze[i] {
			if maze[i][j] == 's' {
				sy = i
				sx = j
			}
			if maze[i][j] == 'g' {
				gy = i
				gx = j
			}
		}
	}
	// fmt.Println(reached[gy][gx])
	// fmt.Println("------")
	dfs(sx, sy)
	// fmt.Println(reached[gy][gx])
	if reached[gy][gx] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
