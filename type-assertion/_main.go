package main

import (
	"fmt"
	"strings"
)

func printIf(src interface{}) {
	switch value := src.(type) {
	case int:
		fmt.Printf("parameter is integer. [value: %d]\n", value)
	case string:
		value = strings.ToUpper(value) // 対象がstring型なのでstringを引数に取る関数が実行できる
		fmt.Printf("parameter is string. [value: %s]\n", value)
	case []string:
		value = append(value, "<不明>") // 対象がsliceなのでAppendができる
		fmt.Printf("parameter is slice string. [value: %s]\n", value)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", src)
	}
}

func main() {
	// printIf(12)
	// printIf("hello")
	printIf([]string{"cat", "dog"})
	// printIf([2]string{"hello", "world"})
}
