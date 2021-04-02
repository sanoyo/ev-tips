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
}
