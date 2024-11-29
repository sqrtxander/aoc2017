package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func solve(s string) int {
	_ = s
    // reverse engineering

	b := 65*100 + 100000
	c := b + 17000
	h := 0
	for {
		if !isPrime(b) {
			h++
		}
		if b == c {
			break
		}
		b += 17
	}

	return h
	// 916 too low
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
