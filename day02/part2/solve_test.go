package main

import "testing"

var INPUT string = `
5 9 2 8
9 4 7 3
3 8 6 5
`[1:]

var EXPECTED int = 9

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
