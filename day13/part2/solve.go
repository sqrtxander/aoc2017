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

func isCaught(time int, modder int) bool {
	return time%modder == 0
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	modders := map[int]int{}

	tMax := 0
	for _, line := range lines {
		t, d, _ := strings.Cut(line, ": ")
		time := utils.HandledAtoi(t)
		depth := utils.HandledAtoi(d)
		modders[time] = 2*depth - 2
		tMax = time
	}

	for i := 0; ; i++ {
		solved := true
		for t := range tMax + 1 {
			if _, ok := modders[t]; !ok {
				continue
			}
			if isCaught(t+i, modders[t]) {
				solved = false
				break
			}
		}
		if solved {
			return i
		}
	}
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
