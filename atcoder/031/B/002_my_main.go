package main

import (
	"fmt"
)

var seamap0 [][]byte
var seamap1 [][]byte

var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

const h = 10
const w = 10

// 条件文
// 再帰関数
func dfs(y, x int) {
	seamap1[y][x] = 'x'
	// 4方向
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if 0 <= nx && nx < 10 && 0 <= ny && ny < 10 && seamap1[ny][nx] == 'o' {
			dfs(nx, ny)
		}
	}
}

func check(map1 [][]byte) bool {
	for i := 0; i < h; i++ {
		for j := 0; j < h; j++ {
			if map1[i][j] == 'o' {
				return false
			}
		}
	}

	return true
}

func main() {
	seamap0 = make([][]byte, h)
	seamap1 = make([][]byte, h)

	for i := range seamap0 {
		seamap0[i] = make([]byte, h)
		seamap1[i] = make([]byte, h)
		fmt.Scan(&seamap0[i])
	}

	// y軸
	for y := 0; y < h; y++ {
		// x軸
		for x := 0; x < w; x++ {
			// oならスキップ
			if seamap0[y][x] == 'o' {
				continue
			}
			seamap1[y][x] = 'o'
			dfs(y, x)
			if check(seamap1) {
				fmt.Println("Yes")
				return
			}
		}
	}
	// 探索が終わっても出力されない場合は、Noを返す
	fmt.Println("No")
}
