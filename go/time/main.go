package main

import (
	"fmt"
	"time"
)

// ref: https://qiita.com/taizo/items/acbee530bd33c803dab4
func main() {
	sample := time.Now()
	// 10分前
	sample = sample.Add(-10 * time.Minute)
	fmt.Println(sample)

	t1 := time.Date(2020, 10, 16, 23, 6, 40, 0, time.Local)
	t2 := time.Date(2019, 10, 17, 23, 6, 40, 0, time.Local)
	// t2よりもt1が以前の時刻かどうか
	fmt.Println(t2.Before(t1))
	// t2よりもt1が先の時刻かどうか
	fmt.Println(t2.After(t1))
}
