package main

import "testing"

var INPUT string = `
99
`[1:]

var EXPECTED int = 8

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
