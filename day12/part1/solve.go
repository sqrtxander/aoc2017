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

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	communicates := map[int][]int{}

	for _, line := range lines {
		id, direct, _ := strings.Cut(line, " <-> ")
		communicates[utils.HandledAtoi(id)] = utils.Map(strings.Split(direct, ", "), utils.HandledAtoi)
	}

	seen := map[int]bool{0: true}
	queue := utils.Queue[int]{0}

	for len(queue) > 0 {
		var q int
		queue, q = queue.Pop()
		for _, direct := range communicates[q] {
			if seen[direct] {
				continue
			}
			seen[direct] = true
			queue = queue.Push(direct)
		}
	}
	return len(seen)
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
