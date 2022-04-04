package main
 
import (
	"fmt"
	"strings"
)

func replace(wk string) string {
	var res = ""
 
	wk = reverse(wk)
	// 全ての文字列に対して、該当する文字列を空に置き換える
	wk = strings.Replace(wk, "resare", "", -1)
	wk = strings.Replace(wk, "esare", "", -1)
	wk = strings.Replace(wk, "remaerd", "", -1)
	wk = strings.Replace(wk, "maerd", "", -1)

	if len(wk) == 0 {
		res = "YES"
	} else {
		res = "NO"
	}
	return res
}
 
//reverse stringreverse
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
 
}
 
func main() {
	var wk string
	fmt.Scan(&wk)
	var res = replace(wk)
	fmt.Println(res)
}
 