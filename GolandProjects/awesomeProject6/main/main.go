package main

import "testing"

func add(a, b int) int {
	return a + b
}
func Testadd(t *testing.T) {
	c := add(1, 2)
	r = 1
	if c != r {
		t.Errorf(c, r)
	}
}
