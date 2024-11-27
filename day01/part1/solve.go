package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	s += string(s[0])
	total := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			num, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatalf("Invlaid number: '%c'\n", s[i])
			}
			total += num
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
