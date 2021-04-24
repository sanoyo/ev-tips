package main

import (
	"testing"
)

func Plus(x, y int) int {
	return x + y
}

func TestPlus(t *testing.T) {
	x, y := 1, 1
	expectation := 3

	res := Plus(x, y)
	if res != expectation {
		t.Errorf("Not equal value %v", res)
	}
}
