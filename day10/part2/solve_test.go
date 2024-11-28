package main

import "testing"

var INPUT string = `
AoC 2017
`[1:]

var EXPECTED string = "33efeb34ea91902bb2f59c9920caa6cd"

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}
