package main

import "testing"

var EXPECTEDS map[string]int = map[string]int{
	"ne,ne,ne":       3,
	"ne,ne,sw,sw":    0,
	"ne,ne,s,s":      2,
	"se,sw,se,sw,sw": 3,
}

func TestSolve(t *testing.T) {
	for input, expected := range EXPECTEDS {
		actual := solve(input)
		if actual != expected {
			t.Fatalf("Expected %d got %d\n", expected, actual)
		}
	}
}
