package main

import "testing"

var INPUT string = `
snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`[1:]

var EXPECTED int = 3

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
