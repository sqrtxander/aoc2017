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
	moves := strings.Split(s, ",")

	p := utils.ORIGIN()
	furthest := 0
	for _, move := range moves {
		switch move {
		case "n":
			p.Y--
		case "ne":
			p.X++
			p.Y--
		case "se":
			p.X++
		case "s":
			p.Y++
		case "sw":
			p.X--
			p.Y++
		case "nw":
			p.X--
		default:
			log.Fatalf("Invalid move: %q\n", move)
		}
		dist := (utils.Abs(p.X) + utils.Abs(p.Y) + utils.Abs(-p.X-p.Y)) / 2
		furthest = max(furthest, dist)
	}

	return furthest
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
