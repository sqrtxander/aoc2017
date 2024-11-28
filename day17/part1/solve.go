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

	buffer := utils.Queue[int]{0}

	for i := 1; i <= 2017; i++ {
		spins := d % len(buffer)
		for range spins {
			var b int
			buffer, b = buffer.Pop()
			buffer = buffer.Push(b)
		}
		buffer = buffer.Push(i)
	}

	return buffer[0]
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
