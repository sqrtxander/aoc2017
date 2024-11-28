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

type linkedList struct {
	val  int
	next *linkedList
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	d := utils.HandledAtoi(s)

	current := -1
	pos := 0
	size := 1
	for i := 1; i <= 50000000; i++ {
		pos = (pos + d) % size
		if pos == 0 {
			current = i
		}
		pos++
		size++
	}

	return current
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
