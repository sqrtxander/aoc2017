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

	stack := utils.Stack[byte]{}
	totalScore := 0
	ip := 0
	for ip < len(s) {
		char := s[ip]
		switch char {
		case '<':
			stack = stack.Push(char)
		case '>':
			firstIdx := -1
			for i := 0; i < len(stack) && firstIdx == -1; i++ {
				if stack[i] == '<' {
					firstIdx = i
				}
			}
			if firstIdx == -1 {
				log.Fatalf("Invalid bracket sequence at i %d\n", ip)
			}
			totalScore += len(stack) - firstIdx - 1
			stack = stack[:firstIdx]
		case '!':
			ip++
		default:
			stack = stack.Push(char)
		}
		ip++
	}

	return totalScore
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

// {{<},{<},{<},{<a>}}
