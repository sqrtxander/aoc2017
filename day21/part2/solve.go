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

func parseSlashLine(line string) [][]byte {
	lines := strings.Split(line, "/")
	result := utils.Map(lines, func(l string) []byte {
		return []byte(l)
	})
	return result
}

func flip(s *[][]byte) {
	slices.Reverse(*s)
}

func transpose(s *[][]byte) {
	for i := range *s {
		for j := range i {
			(*s)[i][j], (*s)[j][i] = (*s)[j][i], (*s)[i][j]
		}
	}
}

func rotate(s *[][]byte) {
	transpose(s)
	flip(s)
}

func hashState(s [][]byte) string {
	result := ""
	for _, line := range s {
		for _, ch := range line {
			result += string(ch)
		}
	}
	return result
}

func generateReplacements(replacements *map[string][][]byte, src [][]byte, dst [][]byte) {
	for range 2 {
		for range 4 {
			(*replacements)[hashState(src)] = dst
			rotate(&src)
		}
		flip(&src)
	}
}

func nextIter(replacements map[string][][]byte, s [][]byte, l int) [][]byte {
	result := make([][]byte, 0, len(s)*(l+1)/l)
	for by := 0; by < len(s); by += l {
		for range l + 1 {
			result = append(result, make([]byte, 0, len(s)*(l+1)/l))
		}
		for bx := 0; bx < len(s); bx += l {
			curr := ""
			for y := by; y < by+l; y++ {
				for x := bx; x < bx+l; x++ {
					curr += string(s[y][x])
				}
			}
			nxt := replacements[curr]
			for y, line := range nxt {
				ydx := len(result) - 1 - l + y
				result[ydx] = append(result[ydx], line...)
			}
		}
	}
	return result
}

func solve(s string, iterations int) int {
	// some optimisations had to be made and unfortunately
	// I had to ditch the cool state representation
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	replacements := map[string][][]byte{}
	for _, line := range lines {
		left, right, _ := strings.Cut(line, " => ")
		generateReplacements(
			&replacements,
			parseSlashLine(left),
			parseSlashLine(right),
		)
	}

	currState := parseSlashLine(".#./..#/###")
	for range iterations {
		for l := 2; l <= 3; l++ {
			if len(currState)%l == 0 {
				currState = nextIter(replacements, currState, l)
				break
			}
		}
	}
	count := 0
	for _, line := range currState {
		for _, ch := range line {
			if ch == '#' {
				count++
			}
		}
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
	fmt.Println(solve(string(contents), 18))
}
