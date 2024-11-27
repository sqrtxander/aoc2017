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

func solve(s string) int {
	s = strings.TrimSpace(s)
	banks := utils.Map(strings.Fields(s), utils.HandledAtoi)

	seen := map[string]bool{}
	count := 0
	for !seen[fmt.Sprint(banks)] {
		seen[fmt.Sprint(banks)] = true
		maxB := slices.Max(banks)
		maxIdx := slices.Index(banks, maxB)
		banks[maxIdx] = 0
		div := maxB / len(banks)
		mod := maxB % len(banks)
		for i := range banks {
			banks[i] += div
		}
		for i := (maxIdx + 1) % len(banks); mod > 0; i = (i + 1) % len(banks) {
			banks[i]++
			mod--
		}
		count++
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
