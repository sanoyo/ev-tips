package main

import (
	"fmt"
)

var seamap1 [][]byte
var seamap2 [][]byte

var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

// 四方向を探索
func dfs(y, x int) {
	for i := 0; i < len(dy); i++ {
		ny := y + dy[i]
		nx := x + dx[i]
		// 条件文
		if 0 <= nx && nx < 10 && 0 <= ny && ny < 10 && seamap2[ny][nx] == 'o' {
			// 再帰関数として設置
			dfs(ny, nx)
		}
	}
}

func check(map1 [][]byte) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if map1[i][j] == 'o' {
				return false
			}
		}
	}
	return true
}

func main() {
	seamap1 := make([][]byte, 10)
	seamap2 := make([][]byte, 10)

	// 入力値を配列に入れる
	for i := range seamap1 {
		seamap1[i] = make([]byte, 10)
		seamap2[i] = make([]byte, 10)
		fmt.Scan(&seamap1[i])
	}

	// xがoにした場合、四方向を探索するためのfor分
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if seamap1[y][x] == 'x' {
				seamap2[y][x] = 'o'
				dfs(y, x)
				if check(seamap2) {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
