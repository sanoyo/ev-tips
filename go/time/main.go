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

	// 時間の切り捨て
	// ref: https://qiita.com/tchnkmr/items/f3c94abb3e3a47e993ab
	t := time.Date(2001, 2, 3, 4, 5, 6, 789000000, time.UTC)
	fmt.Println(t)                          // -> 2001-02-03 04:05:06.789 +0000 UTC Truncate前
	fmt.Println(t.Truncate(time.Hour * 24)) // -> 2001-02-03 00:00:00 +0000 UTC
	fmt.Println(t.Truncate(time.Hour))      // -> 2001-02-03 04:00:00 +0000 UTC
	fmt.Println(t.Truncate(time.Minute))    // -> 2001-02-03 04:05:00 +0000 UTC
	fmt.Println(t.Truncate(time.Second))    // -> 2001-02-03 04:05:06 +0000 UTC
	t = time.Date(2001, 2, 3, 4, 5, 6, 789000000, time.Local)
	fmt.Println(t.Truncate(time.Hour * 24)) // -
}
