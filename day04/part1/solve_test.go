package main

import "testing"

var INPUT string = `
aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa
`[1:]

var EXPECTED int = 2

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
