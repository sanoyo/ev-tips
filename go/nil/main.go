package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)
	// s34 = nil

	fmt.Println("s1", len(s1), cap(s1), s1 == nil, s1[:], s1[:] == nil)
	fmt.Println("s2", len(s2), cap(s2), s2 == nil, s2[:], s2[:] == nil)
	fmt.Println("s3", len(s3), cap(s3), s3 == nil, s3[:], s3[:] == nil)

	for range s1 {
	}
	for range s2 {
	}
	for range s3 {
	}
}
