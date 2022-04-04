package main

import (
	"fmt"
)

// // 実装順序
// // 左と右の文字列を分解する
// // この時どちらも配列で扱えるようにしておく
// // 最初に入力された数値でfor文を回す
// // 交互に文字列を出力する

func main() {
	var number int
	var text, text2 string

	fmt.Scan(&number, &text, &text2)

	var ans string
	for i := 0; i < number; i++ {
		ans += string(text[i])
		ans += string(text2[i])
	}

	fmt.Println(ans)
}
