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

	bhgrid := utils.ParseBoundedHashGrid(s, '.', '#')
	pos := utils.Point{X: bhgrid.W / 2, Y: bhgrid.H / 2}
	dir := utils.UP
	infected := bhgrid.Grid

	count := 0
	for range 10000 {
		if infected[pos] {
			dir.RotateRight()
		} else {
			dir.RotateLeft()
			count++
		}
		infected[pos] = !infected[pos]
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
