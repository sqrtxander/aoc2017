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

func getContribution(row []int) int {
	for i, n1 := range row {
		for _, n2 := range row[i+1:] {
			if n1%n2 == 0 {
				return n1 / n2
			} else if n2%n1 == 0 {
				return n2 / n1
			}
		}
	}
	return -1
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	numStrs := utils.Map(lines, strings.Fields)
	sheet := utils.Map(numStrs, func(line []string) []int {
		return utils.Map(line, utils.HandledAtoi)
	})
	chsum := 0
	for _, row := range sheet {
		contribution := getContribution(row)
		if contribution == -1 {
			log.Fatalln("Invalid input")
		}
		chsum += contribution
	}

	return chsum
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
