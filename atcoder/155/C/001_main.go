package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
 
func solution(n int, texts []string) []string {
	ans := make([]string, 0)
	var c int
	seen := make(map[string]int, n)
 
	// mapで同じものをカウント
	// map[beat:1 bed:1 beet:2 bet:1 vet:2]
	for _, text := range texts {
		seen[text]++
		c = max(c, seen[text])
	}
 
	for k, v := range seen {
		// value(数値)とmaxのカウント数が同じなのであれば、ansに追加
		if v == c {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	return ans
}
 
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
 
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
 
	sc.Scan()
	// 文字列カウント
	n, _ := strconv.Atoi(sc.Text())
 
	// 文字列を配列に入れる
	texts := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		texts[i] = sc.Text()
	}
 
	// 一番出てきた回数が多いものを残す
	ans := solution(n, texts)
	// 残した配列を順番に出力する
	for _, text := range ans {
		fmt.Println(text)
	}
}
