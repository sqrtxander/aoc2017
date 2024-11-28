package main

import "testing"

var INPUT string = `
Generator A starts with 65
Generator B starts with 8921
`[1:]

var EXPECTED int = 309

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
