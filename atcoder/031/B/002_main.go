// 002_main.go
// https://atcoder.jp/contests/arc031/submissions/7030599
package main

import "fmt"

var seamap0 [][]byte
var seamap1 [][]byte

// var reached [][]bool

const h = 10
const w = 10

var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func dfs(y, x int) {
	//fmt.Println("dfs en: x", x, "y", y)
	seamap1[y][x] = 'x'
	for i := 0; i < 4; i++ {
		// 上だけの場合は、dx=1, dy=0にする必要があるため、このような変数を準備する
		// 横移動
		nx := x + dx[i]
		// 上下移動
		ny := y + dy[i]
		if 0 <= nx && nx < w && 0 <= ny && ny < h && seamap1[ny][nx] == 'o' {
			dfs(ny, nx)
		}
	}
}

func mapinit() {
	for i := range seamap0 {
		copy(seamap1[i], seamap0[i])
	}
}

func check(smap [][]byte) bool {
	fmt.Println(smap)
	for i := range smap {
		for j := range smap[i] {
			if smap[i][j] == 'o' {
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
		seamap0[i] = make([]byte, w)
		seamap1[i] = make([]byte, w)
		fmt.Scan(&seamap0[i])
	}
	for y := range seamap0 {
		for x := range seamap0[y] {
			if seamap0[y][x] == 'x' { // (i,j)を埋め立てた場合、
				//fmt.Printf("[%d, %d] 埋め立てる \n", j, i)

				// seamap1に、seamap0と同じ配列の状態に戻す
				mapinit()
				seamap1[y][x] = 'o'
				dfs(y, x)
				if check(seamap1) {
					fmt.Println("YES")
					// fmt.Printf("YES - (%d,%d)\n", x+1, y+1)
					return
				}
			}
		}
	}
	fmt.Println("NO")
}
