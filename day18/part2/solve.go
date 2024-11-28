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

type program struct {
	regs     map[string]int
	ip       int
	waiting  bool
	finished bool
	inQueue  utils.Queue[int]
	sent     int
}

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

	p0 := &program{
		regs:     map[string]int{"p": 0},
		ip:       0,
		waiting:  false,
		finished: false,
		inQueue:  utils.Queue[int]{},
		sent:     0,
	}
	p1 := &program{
		regs:     map[string]int{"p": 1},
		ip:       0,
		waiting:  false,
		finished: false,
		inQueue:  utils.Queue[int]{},
		sent:     0,
	}
	running := p0
	waiting := p1
	for !p0.finished || !p1.finished {
		if running.finished {
			running, waiting = waiting, running
		}
		inst := instructions[running.ip]
		switch inst[0] {
		case "snd":
			waiting.inQueue = waiting.inQueue.Push(getRegOrValue(inst[1], running.regs))
			running.sent++
			running.ip++
		case "set":
			running.regs[inst[1]] = getRegOrValue(inst[2], running.regs)
			running.ip++
		case "add":
			running.regs[inst[1]] += getRegOrValue(inst[2], running.regs)
			running.ip++
		case "mul":
			running.regs[inst[1]] *= getRegOrValue(inst[2], running.regs)
			running.ip++
		case "mod":
			running.regs[inst[1]] %= getRegOrValue(inst[2], running.regs)
			running.ip++
		case "rcv":
			if len(running.inQueue) == 0 {
				if (waiting.waiting && len(waiting.inQueue) == 0) || waiting.finished {
					running.finished = true
					waiting.finished = true
					break
				}
				running.waiting = true
				waiting.waiting = false
				running, waiting = waiting, running
			} else {
				running.inQueue, running.regs[inst[1]] = running.inQueue.Pop()
				running.ip++
			}
		case "jgz":
			if getRegOrValue(inst[1], running.regs) > 0 {
				running.ip += getRegOrValue(inst[2], running.regs)
			} else {
				running.ip++
			}
		default:
			log.Fatalf("Invalid instruction: %q\n", inst[0])
		}
		if running.ip < 0 || running.ip >= len(instructions) {
			running.finished = true
		}
	}

	return p1.sent
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
