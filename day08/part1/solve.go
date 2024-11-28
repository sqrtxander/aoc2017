package main

import (
	"aoc2017/utils"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	instructions := utils.Map(lines, strings.Fields)

	regs := map[string]int{}
	for _, inst := range instructions {
		reg := inst[0]
		val := utils.HandledAtoi(inst[2])
		if inst[1] == "dec" {
			val *= -1
		} else if inst[1] != "inc" {
			log.Fatalf("Invalid instruction: %q\n", inst[1])
		}
		conditionalReg := inst[4]
		conditionalVal := utils.HandledAtoi(inst[6])
		var conditionMet bool
		switch inst[5] {
		case ">":
			conditionMet = regs[conditionalReg] > conditionalVal
		case "<":
			conditionMet = regs[conditionalReg] < conditionalVal
		case ">=":
			conditionMet = regs[conditionalReg] >= conditionalVal
		case "<=":
			conditionMet = regs[conditionalReg] <= conditionalVal
		case "==":
			conditionMet = regs[conditionalReg] == conditionalVal
		case "!=":
			conditionMet = regs[conditionalReg] != conditionalVal
		default:
			log.Fatalf("Invalid operator: %q\n", inst[5])
		}
		if conditionMet {
			regs[reg] += val
		}
	}
	maxVal := math.MinInt
	for _, v := range regs {
		maxVal = max(maxVal, v)
	}
	return maxVal
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
