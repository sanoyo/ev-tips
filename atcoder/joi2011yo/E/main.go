// https://atcoder.jp/contests/joi2011yo/submissions/7035859
package main

import "fmt"

type q struct {
	data []point
}

type point struct {
	x, y int
}

func (q *q) q(v point) {
	q.data = append(q.data, v)
}

func (q *q) dq(v *point) {
	if len(q.data) > 0 {
		*v = q.data[0]
		q.data = q.data[1:]
	}
}

var h, w, n int

// var s, g point
var maze [][]byte
var d [][]int
var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func bfs(s, g *point) int {
	//fmt.Printf("bfs en: s(%d,%d), g(%d,%d)\n", s.x+1, s.y+1, g.x+1, g.y+1)
	var q q
	d[s.y][s.x] = 0
	q.q(point{s.x, s.y})
	//fmt.Println("len(q), q.data", len(q.data), q.data)
	for len(q.data) > 0 {
		var p point
		q.dq(&p)
		//fmt.Printf("(%d,%d) dq\n", p.x+1, p.y+1)
		for i := 0; i < 4; i++ {
			nx := p.x + dx[i]
			ny := p.y + dy[i]
			if 0 <= ny && ny < h && 0 <= nx && nx < w && d[ny][nx] == -1 && maze[ny][nx] != 'X' {
				d[ny][nx] = d[p.y][p.x] + 1
				q.q(point{nx, ny})
			}
		}
	}
	return d[g.y][g.x]
}

// 配列を全て-1にする
func dinit() {
	for i := range d {
		for j := range d[i] {
			d[i][j] = -1
		}
	}
}

func main() {
	fmt.Scan(&h, &w, &n)
	var s point
	var cheese []point
	maze = make([][]byte, h)
	d = make([][]int, h)
	cheese = make([]point, n)
	for i := range maze {
		d[i] = make([]int, w)
		maze[i] = make([]byte, w)
		fmt.Scan(&maze[i])
		for j := range maze[i] {
			if maze[i][j] == 'S' {
				s.x = j
				s.y = i
			}
			if '0' <= maze[i][j] && maze[i][j] <= '9' {
				cheese[maze[i][j]-'0'-1].x = j
				cheese[maze[i][j]-'0'-1].y = i
			}
		}
	}
	//fmt.Println("h", h, "w", w, "n", n, "maze", maze, "s", s, "cheese", cheese)
	ans := 0
	for i := range cheese {
		var g point
		g.y = cheese[i].y
		g.x = cheese[i].x
		dinit()
		ans += bfs(&s, &g)
		s = g
	}
	fmt.Println(ans)
}
