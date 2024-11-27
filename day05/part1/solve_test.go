package main

import "testing"

var INPUT string = `
0
3
0
1
-3
`[1:]

var EXPECTED int = 5

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
