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

type state struct {
	state map[utils.Point]bool
	size  int
}

func parseSlashLine(line string) state {
	lines := strings.Split(line, "/")
	result := map[utils.Point]bool{}
	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case '#':
				result[utils.Point{X: x, Y: y}] = true
			case '.':
				break
			default:
				log.Fatalf("Invalid character: %q\n", ch)
			}
		}
	}
	return state{
		state: result,
		size:  len(lines),
	}
}

func deParseSlashLine(s state) string {
	result := ""
	for y := range s.size {
		result += "/"
		for x := range s.size {
			if s.state[utils.Point{X: x, Y: y}] {
				result += "#"
			} else {
				result += "."
			}
		}
	}
	return result[1:]
}

func flip(s state) state {
	result := map[utils.Point]bool{}
	for y := range s.size {
		for x := range s.size {
			if s.state[utils.Point{X: x, Y: s.size - y - 1}] {
				result[utils.Point{X: x, Y: y}] = true
			}
		}
	}
	return state{
		state: result,
		size:  s.size,
	}
}

func rotate(s state) state {
	result := map[utils.Point]bool{}
	for y := range s.size {
		for x := range s.size {
			if s.state[utils.Point{X: s.size - y - 1, Y: x}] {
				result[utils.Point{X: x, Y: y}] = true
			}
		}
	}
	return state{
		state: result,
		size:  s.size,
	}
}

func hashState(s state) string {
	// return deParseSlashLine(s)
	result := ""
	for y := range s.size {
		for x := range s.size {
			if s.state[utils.Point{X: x, Y: y}] {
				result += "#"
			} else {
				result += "."
			}
		}
	}
	return result
}

func generateReplacements(replacements *map[string]state, src state, dst state) {
	for range 2 {
		for range 4 {
			(*replacements)[hashState(src)] = dst
			src = rotate(src)
		}
		src = flip(src)
	}
}

func toChoppedStates(s state, l int) [][]string {
	result := make([][]string, 0, s.size*(l+1)/l)
	for by := 0; by < s.size; by += l {
		result = append(result, make([]string, 0, s.size*(l+1)/l))
		for bx := 0; bx < s.size; bx += l {
			curr := ""
			for y := by; y < by+l; y++ {
				for x := bx; x < bx+l; x++ {
					if s.state[utils.Point{X: x, Y: y}] {
						curr += "#"
					} else {
						curr += "."
					}
				}
			}
			idx := len(result) - 1
			result[idx] = append(result[idx], curr)
		}
	}
	return result
}

func repair(sss [][]state) state {
	result := state{
		state: map[utils.Point]bool{},
		size:  sss[0][0].size * len(sss),
	}
	for by, ss := range sss {
		for bx, s := range ss {
			for p := range s.state {
				result.state[utils.Point{
					X: p.X + s.size*bx,
					Y: p.Y + s.size*by,
				}] = true
			}
		}
	}
	return result
}

func printState(s state) {
	bhg := utils.BoundedHashGrid{
		Grid: s.state,
		W:    s.size,
		H:    s.size,
	}
	fmt.Println(bhg.GetBoundedHash())
}

func solve(s string, iterations int) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	replacements := map[string]state{}
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
		// printState(currState)
		for l := 2; l <= 3; l++ {
			if currState.size%l == 0 {
				chopped := toChoppedStates(currState, l)
				larger := utils.Map(chopped, func(ss []string) []state {
					return utils.Map(ss, func(s string) state {
						return replacements[s]
					})
				})
				currState = repair(larger)
				break
			}
		}
	}
	return len(currState.state)
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
	fmt.Println(solve(string(contents), 5))
}
