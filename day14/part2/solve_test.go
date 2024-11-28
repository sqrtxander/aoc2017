package main

import "testing"

var INPUT string = `
flqrgnkx
`[1:]

var EXPECTED int = 1242

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
