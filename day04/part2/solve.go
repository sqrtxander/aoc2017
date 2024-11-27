package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func isValid(line string) bool {
	seen := map[string]bool{}
	for _, word := range strings.Fields(line) {
		bytes := []byte(word)
		slices.Sort(bytes)
		word = string(bytes)
		if _, ok := seen[word]; ok {
			return false
		}
		seen[word] = true
	}
	return true
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	total := 0
	for _, line := range lines {
		if isValid(line) {
			total++
		}
	}

	return total
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
