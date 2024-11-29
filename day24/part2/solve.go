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

type bridge struct {
	ports     []*utils.Pair[int, int]
	nextMatch int
	strength  int
}

func parsePort(line string) *utils.Pair[int, int] {
	l, r, _ := strings.Cut(line, "/")
	return &utils.Pair[int, int]{
		K: utils.HandledAtoi(l),
		V: utils.HandledAtoi(r),
	}
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	allPorts := utils.Map(lines, parsePort)
	queue := utils.Queue[*bridge]{}
	for _, port := range allPorts {
		if port.K == 0 {
			queue = queue.Push(&bridge{
				ports:     []*utils.Pair[int, int]{port},
				nextMatch: port.V,
				strength:  port.V,
			})
		} else if port.V == 0 {
			queue = queue.Push(&bridge{
				ports:     []*utils.Pair[int, int]{port},
				nextMatch: port.K,
				strength:  port.K,
			})
		}
	}
	maxStrength := 0
	maxLen := 0
	for len(queue) > 0 {
		var b *bridge
		queue, b = queue.Pop()
		if len(b.ports) == maxLen {
			maxStrength = max(maxStrength, b.strength)
		} else if len(b.ports) > maxLen {
			maxStrength = b.strength
			maxLen = len(b.ports)
		}
		match := false
		for _, port := range allPorts {
			if slices.Contains(b.ports, port) {
				continue
			}
			if port.K != b.nextMatch && port.V != b.nextMatch {
				continue
			}
			if match {
				_ = match
			}
			match = true
			nextMatch := port.K
			if port.K == b.nextMatch {
				nextMatch = port.V
			}
			newPorts := append(slices.Clone(b.ports), port)
			queue = queue.Push(&bridge{
				ports:     newPorts,
				nextMatch: nextMatch,
				strength:  b.strength + port.K + port.V,
			})
		}
	}

	return maxStrength
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
