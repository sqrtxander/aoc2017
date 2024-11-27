package main

import "testing"

var INPUT string = `
0  2  7  0
`[1:]

var EXPECTED int = 4

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
