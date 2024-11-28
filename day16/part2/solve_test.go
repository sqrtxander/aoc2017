package main

import "testing"

var INPUT string = `
s1,x3/4,pe/b
`[1:]

var EXPECTED string = "abcde"

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 5)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}
