package main

import (
	"fmt"
)

var sx, sy int       // スタートの座標
var gx, gy int       // ゴールの座標
var reached [][]bool // 座標(x,y)に到達したか
var maze [][]byte    // 迷路の入力を受け取る配列()

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, -1, 1}

func dfs(x byte, y byte) {

}

func main() {
	maze := make([][]byte, 10)
	reached := make([][]bool, 10)
	fmt.Println(maze)

	for i := range maze {
		maze[i] = make([]byte, 10)
		reached[i] = make([]bool, 10)

		fmt.Scan(&maze[i])
		for j := range maze[i] {
			if maze[i][j] == 'o' {
				sy = i
				sx = j
			} else {
				gx = i
				gy = j
			}
		}

		dfs(sx, sy)
		fmt.Println("Yes")
		fmt.Println("No")
	}
}
