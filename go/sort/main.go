// https://qiita.com/Sekky0905/items/2d5ccd6d076106e9d21c
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "V", Age: 3},
		{Name: "K", Age: 3},
		{Name: "Y", Age: 3},
		{Name: "A", Age: 4},
		{Name: "E", Age: 3},
		{Name: "D", Age: 1},
		{Name: "C", Age: 3},
		{Name: "X", Age: 2},
		{Name: "B", Age: 3},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("NameでSort(Not-Stable):%+v\n", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("AgeでSort(Not-Stable):%+v\n", people)

	people2 := []Person{
		{Name: "V", Age: 3},
		{Name: "K", Age: 3},
		{Name: "Y", Age: 3},
		{Name: "A", Age: 4},
		{Name: "E", Age: 3},
		{Name: "D", Age: 1},
		{Name: "C", Age: 3},
		{Name: "X", Age: 2},
		{Name: "B", Age: 3},
	}

	sort.SliceStable(people2, func(i, j int) bool {
		return people2[i].Name < people2[j].Name
	})
	fmt.Printf("NameでSort(Stable):%+v\n", people2)

	sort.SliceStable(people2, func(i, j int) bool {
		return people2[i].Age < people2[j].Age
	})
	fmt.Printf("AgeでSort:%+v\n", people2)
}
