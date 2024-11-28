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
	currentContributor := 0
	totalScore := 0
	ip := 0
	inGarbage := false
	for ip < len(s) {
		char := s[ip]
		switch char {
		case '{':
			if !inGarbage {
				currentContributor += 1
				stack = stack.Push(char)
			}
		case '}':
			if !inGarbage && stack.Peek() != '{' {
				log.Fatalf("Invalid bracket sequence at idx %d\n", ip)
			}
			if !inGarbage {
				stack, _ = stack.Pop()
				totalScore += currentContributor
				currentContributor--
			}
		case '<':
			stack = stack.Push(char)
			inGarbage = true
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
			stack = stack[:firstIdx]
			inGarbage = false
		case '!':
			ip++
		default:
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
