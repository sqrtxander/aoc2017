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
	d := utils.HandledAtoi(s)

	p := utils.ORIGIN()
	dir := utils.RIGHT

	grid := map[utils.Point]int{p: 1}

	for grid[p] < d {
		if p.X == p.Y && p.X > 0 {
			dir.RotateRight()
		}
		if -p.X == p.Y && p.X < 0 {
			dir.RotateRight()
		}
		if p.X == p.Y && p.X < 0 {
			dir.RotateRight()
		}
		if p.X == -p.Y+1 && p.X > 0 {
			dir.RotateRight()
		}
		p.MoveInDir(dir, 1)
		newVal := 0
		for _, n := range utils.Adjacent8(p) {
			if val, ok := grid[n]; ok {
				newVal += val
			}
		}
		grid[p] = newVal
	}
	return grid[p]
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
