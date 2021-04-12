package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "太郎",
		Age:  20,
	}

	p2 := Person{
		Name: "太郎",
		Age:  20,
	}

	// ref : https://qiita.com/Sekky0905/items/1ff4979d80b163e0aeb6
	fmt.Println("単純な構造体の比較")
	fmt.Printf("p1 == p2 : 等価(メモリが別でも値が一緒だったらOK) : %t\n", p1 == p2)
	fmt.Printf("&p1 == &p2 : 等値 : %t\n", &p1 == &p2)
	fmt.Printf("reflect.DeepEqual(p1, p2) : 等価 : %t\n", reflect.DeepEqual(p1, p2))
	fmt.Printf("reflect.DeepEqual(&p1, &p2)  : %t\n", reflect.DeepEqual(&p1, &p2))

	fmt.Println("単純な構造体の比較2")
	p3 := p1
	fmt.Printf("p1 == p3 : 等価(メモリが別でも値が一緒だったらOK) : %t\n", p1 == p3)
	fmt.Printf("&p1 == &p3 : 等値 : %t\n", &p1 == &p3)
	fmt.Printf("reflect.DeepEqual(p1, p3) : 等価 : %t\n", reflect.DeepEqual(p1, p3))
	fmt.Printf("reflect.DeepEqual(&p1, &p3) : %t\n", reflect.DeepEqual(&p1, &p3))
}
