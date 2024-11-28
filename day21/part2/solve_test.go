package main

import "testing"

var INPUT string = `
../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#
`[1:]

var EXPECTED int = 12

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 2)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
