package main

import (
	"fmt"
	"reflect"
)

type T struct{}
type U struct{}

func main() {
	var b *T
	var c = (*U)(b)
	fmt.Println(c == nil)
	fmt.Println(reflect.TypeOf(c))
}
