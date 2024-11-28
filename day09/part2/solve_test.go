package main

import "testing"

var INPUT string = `
<{o"i!a,<{i<a>
`[1:]

var EXPECTED int = 10

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
