package main

import "fmt"

// ref: https://simple-minds-think-alike.moritamorie.com/entry/slice-operations
func main() {
	//  スライスを作る
	people := []string{
		"竈門 炭治郎(かまど たんじろう)",
		"我妻 善逸(あがつま ぜんいつ）",
		"嘴平 伊之助(はしびら いのすけ)",
	}
	// // 別のスライスに代入
	// people2 := people

	// // 2つのスライスが指す配列のポインタは同じ値になる。
	// fmt.Println(&people[0])
	// fmt.Println(&people2[0])

	// -------------------------------------------
	slice := []string{"apple", "banana", "peace"}
	// 添字1〜添字3-1(=2)までの要素で新しくスライスを作る
	partical := slice[1:3]
	fmt.Println(partical)
	// 添字1〜添字3-1(=2)までの要素で新しくスライスを作る
	people2 := slice[1:3]

	// 元のスライスの2番目の要素のポインタと
	// 新しくできたスライスの1番目の要素のポインタは同じ値になる。
	fmt.Println(&people[1])
	fmt.Println(&people2[0])

	// -------------------------------------------
	// before
	// これも同じになる
	// ids := []int64{1, 2, 3}
	// for _, id := range ids {
	// 	fmt.Println(&id)
	// }
	// 0xc0000140c8
	// 0xc0000140c8
	// 0xc0000140c8

	// after
	// for _, id := range ids {
	// 	replaceID := id
	// 	fmt.Println(&replaceID)
	// }
	// 0xc0000140e0
	// 0xc0000140e8
	// 0xc0000140f0
}
