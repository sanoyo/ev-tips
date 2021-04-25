// https://tutuz-tech.hatenablog.com/entry/2019/10/20/145302
package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)

	fmt.Println("s1", len(s1), cap(s1), s1 == nil, s1[:], s1[:] == nil)
	fmt.Println("s2", len(s2), cap(s2), s2 == nil, s2[:], s2[:] == nil)
	fmt.Println("s3", len(s3), cap(s3), s3 == nil, s3[:], s3[:] == nil)
}
