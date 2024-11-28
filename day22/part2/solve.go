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

const (
	clean int = iota
	weakened
	infected
	flagged
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	states := map[utils.Point]int{}
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '.':
				states[utils.Point{X: x, Y: y}] = clean
			case '#':
				states[utils.Point{X: x, Y: y}] = infected
			default:
				log.Fatalf("Invalid hashdot character: '%c'\n", char)
			}
		}
	}

	pos := utils.Point{X: len(lines[0]) / 2, Y: len(lines) / 2}
	dir := utils.UP

	count := 0
	for range 10000000 {
		switch states[pos] {
		case clean:
			dir.RotateLeft()
		case weakened:
			count++
		case infected:
			dir.RotateRight()
		case flagged:
			dir.Rotate180()
		default:
			log.Fatalln("Invalid state")
		}
		states[pos] = (states[pos] + 1) % 4
		pos.MoveInDir(dir, 1)
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
