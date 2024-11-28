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

func genNext(prev int, mult int) int {
	return (prev * mult) % 2147483647
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	valA := utils.HandledAtoi(strings.Fields(lines[0])[4])
	valB := utils.HandledAtoi(strings.Fields(lines[1])[4])

	count := 0
	for range 40000000 {
		valA = genNext(valA, 16807)
		valB = genNext(valB, 48271)
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
