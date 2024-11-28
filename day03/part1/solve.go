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

	l := 0
	r := d
	var curr int
	for l <= r {
		m := (l + r) / 2
		if m*m == d {
			curr = m
			break
		} else if m*m <= d {
			curr = m
			l = m + 1
		} else if m*m > d {
			r = m - 1
		}
	}
	var p utils.Point
	var dir utils.Direction
	if curr%2 == 0 {
		p = utils.Point{
			X: -curr/2 + 1,
			Y: curr / 2,
		}
		dir = utils.LEFT
	} else {
		p = utils.Point{
			X: curr / 2,
			Y: -curr / 2,
		}
		dir = utils.RIGHT
	}

	curr *= curr
	for curr != d {
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
		curr++
	}
	return p.Manhattan()
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
