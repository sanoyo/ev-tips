package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	text := make(map[int]string, n)
	for key, value := range text {
		map[int]string{key: fmt.Scan(&value)}
	}

	// fmt.Println(text)
}
