package main

import (
	"aoc2017/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func solve(s string) string {
	// s = strings.TrimSpace(s) // :(
	lines := strings.Split(s, "\n")

	var loc utils.Point
	dir := utils.DOWN
	grid := map[utils.Point]rune{}
	for y, line := range lines {
		for x, ch := range line {
			if ch != ' ' {
				grid[utils.Point{X: x, Y: y}] = ch
			}
			if ch != ' ' && y == 0 {
				loc = utils.Point{X: x, Y: y}
			}
		}
	}

	result := ""
	for {
		ch, ok := grid[loc]
		if !ok {
			break
		}
		if !slices.Contains([]rune("|-+"), ch) {
			result += string(ch)
		}
		if ch == '+' && (dir == utils.UP || dir == utils.DOWN) {
			loc.MoveInDir(utils.RIGHT, 1)
			if _, ok := grid[loc]; ok {
				dir = utils.RIGHT
			} else {
				dir = utils.LEFT
			}
			loc.MoveInDir(utils.LEFT, 1)
		} else if ch == '+' && (dir == utils.RIGHT || dir == utils.LEFT) {
			loc.MoveInDir(utils.UP, 1)
			if _, ok := grid[loc]; ok {
				dir = utils.UP
			} else {
				dir = utils.DOWN
			}
			loc.MoveInDir(utils.DOWN, 1)
		}
		loc.MoveInDir(dir, 1)
	}

	return result
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
