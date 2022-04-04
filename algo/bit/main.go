package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// 2^n乗のfor文
	for bit := 0; bit < (1 << uint64(n)); bit++ {
		// 3を入力した場合、0~7の計8つが出力される。
		// fmt.Println(bit)
		for i := 0; i < n; i++ {
			// 3を入力した場合、0, 1, 2のセットが計8つが出力される。
			fmt.Println(i)
			// bitsのi個目の要素の状態がonかどうかチェック
			if (bit>>uint64(i))&1 == 1 {
				fmt.Println("aaa")
				fmt.Println("-----")
			}
		}
	}
}
