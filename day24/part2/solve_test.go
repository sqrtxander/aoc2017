package main

import "testing"

var INPUT string = `
0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
`[1:]

var EXPECTED int = 19

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
