package main

import (
	"aoc2017/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func getRegOrValue(arg string, regs map[string]int) int {
	val, err := strconv.Atoi(arg)
	if err != nil {
		return regs[arg]
	}
	return val
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	instructions := utils.Map(lines, strings.Fields)

	ip := 0
	var sound int
	regs := map[string]int{}
	for ip >= 0 && ip < len(instructions) {
		inst := instructions[ip]
		switch inst[0] {
		case "snd":
			sound = getRegOrValue(inst[1], regs)
			ip++
		case "set":
			regs[inst[1]] = getRegOrValue(inst[2], regs)
			ip++
		case "add":
			regs[inst[1]] += getRegOrValue(inst[2], regs)
			ip++
		case "mul":
			regs[inst[1]] *= getRegOrValue(inst[2], regs)
			ip++
		case "mod":
			regs[inst[1]] %= getRegOrValue(inst[2], regs)
			ip++
		case "rcv":
			if getRegOrValue(inst[1], regs) != 0 {
				return sound
			}
			ip++
		case "jgz":
			if getRegOrValue(inst[1], regs) > 0 {
				ip += getRegOrValue(inst[2], regs)
			} else {
				ip++
			}
		default:
			log.Fatalf("Invalid instruction: %q\n", inst[0])
		}
	}

	return -1
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
