package main

import "testing"

var INPUT string = `
3,4,1,5
`[1:]

var EXPECTED int = 12

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 5)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}
