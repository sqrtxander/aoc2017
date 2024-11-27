package main

import "testing"

var INPUT string = `
5 1 9 5
7 5 3
2 4 6 8
`[1:]

var EXPECTED int = 18

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
