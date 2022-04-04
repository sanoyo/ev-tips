package main

import "fmt"

type Stronger interface {
	Body()
	Brain()
}

type Animal struct {
	height int
}

type Human struct {
	brain string
}

func (h Animal) Body() {
	fmt.Println("体でかい")
}

func (h Human) Brain() {
	fmt.Println("頭脳派")
}

// インターフェースが引数になるメソッド
func Really(s Stronger) {
	s.Body()
	s.Brain()
}

func main() {
	var animal Animal = Animal{height: 300}
	fmt.Println(animal) // {300}
	var wow Stronger
	fmt.Println(wow) // nil
}
