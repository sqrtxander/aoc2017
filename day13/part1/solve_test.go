package main

import "testing"

var INPUT string = `
0: 3
1: 2
4: 4
6: 4
`[1:]

var EXPECTED int = 24

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
