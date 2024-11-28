package main

import (
	"aoc2017/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func genNext(prev int, mult int, div int) int {
	result := (prev * mult) % 2147483647
	for (result % div) != 0 {
		result = (result * mult) % 2147483647
	}
	return result
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	valA := utils.HandledAtoi(strings.Fields(lines[0])[4])
	valB := utils.HandledAtoi(strings.Fields(lines[1])[4])

	count := 0
	for range 5000000 {
		valA = genNext(valA, 16807, 4)
		valB = genNext(valB, 48271, 8)
		if valA&0xffff == valB&0xffff {
			count++
		}
	}

	return count
}

func main() {
	var inputPath string
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	} else {
		_, currentFilePath, _, _ := runtime.Caller(0)
		dir := filepath.Dir(currentFilePath)
		dir = filepath.Dir(dir)
		inputPath = filepath.Join(dir, "input.in")
	}
	contents, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file %s:\n%v\n", inputPath, err)
		return
	}
	fmt.Println(solve(string(contents)))
}
