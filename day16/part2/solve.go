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

func solve(s string, chars int) string {
	s = strings.TrimSpace(s)
	instructions := strings.Split(s, ",")
	programs := utils.Deque[byte]{}
	for i := 0; i < chars; i++ {
		programs = programs.PushRight(byte(i + 0x61))
	}

	store := map[string]int{}

	for i := 0; i < 1000000000; i++ {
		for _, inst := range instructions {
			switch inst[0] {
			case 's':
				x := utils.HandledAtoi(inst[1:])
				for range x {
					var p byte
					programs, p = programs.PopRight()
					programs = programs.PushLeft(p)
				}
			case 'x':
				aStr, bStr, _ := strings.Cut(inst[1:], "/")
				a := utils.HandledAtoi(aStr)
				b := utils.HandledAtoi(bStr)
				programs[a], programs[b] = programs[b], programs[a]
			case 'p':
				a := slices.Index(programs, inst[1])
				b := slices.Index(programs, inst[3])
				programs[a], programs[b] = programs[b], programs[a]
			default:
				log.Fatalf("Invalid instruction: %q", inst)
			}
		}
		if idx, ok := store[string(programs)]; ok {
			cycleLen := i - idx
			left := 1000000000 - i - 1
			i += cycleLen * (left / cycleLen)
		}
		store[string(programs)] = i
	}

	return string(programs)
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
	fmt.Println(solve(string(contents), 16))
}
