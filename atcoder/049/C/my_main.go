package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Println(t)

	de := []string{"maerd", "remaerd", "esare", "reesare"}
	for i := 0; i < len(de); i++ {
		fmt.Println(Reverse(s))
		fmt.Println(strings.Split(t[:], de[i]))
	}
}
