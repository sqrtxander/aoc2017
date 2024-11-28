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
	dir := utils.SOUTH
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
		if ch == '+' && (dir == utils.NORTH || dir == utils.SOUTH) {
			loc.MoveInDir(utils.EAST, 1)
			if _, ok := grid[loc]; ok {
				dir = utils.EAST
			} else {
				dir = utils.WEST
			}
			loc.MoveInDir(utils.WEST, 1)
		} else if ch == '+' && (dir == utils.EAST || dir == utils.WEST) {
			loc.MoveInDir(utils.NORTH, 1)
			if _, ok := grid[loc]; ok {
				dir = utils.NORTH
			} else {
				dir = utils.SOUTH
			}
			loc.MoveInDir(utils.SOUTH, 1)
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
