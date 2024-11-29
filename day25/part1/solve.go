package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	a int = iota
	b
	c
	d
	e
	f
)

func solve(s string) int {
	_ = s

	cursor := 0
	tape := map[int]bool{}
	state := a
	for range 12261543 {
		switch state {
		case a:
			if !tape[cursor] {
				tape[cursor] = true
				cursor++
				state = b
			} else {
				delete(tape, cursor)
				cursor--
				state = c
			}
		case b:
			if !tape[cursor] {
				tape[cursor] = true
				cursor--
				state = a
			} else {
				tape[cursor] = true
				cursor++
				state = c
			}
		case c:
			if !tape[cursor] {
				tape[cursor] = true
				cursor++
				state = a
			} else {
				delete(tape, cursor)
				cursor--
				state = d
			}
		case d:
			if !tape[cursor] {
				tape[cursor] = true
				cursor--
				state = e
			} else {
				tape[cursor] = true
				cursor--
				state = c
			}
		case e:
			if !tape[cursor] {
				tape[cursor] = true
				cursor++
				state = f
			} else {
				tape[cursor] = true
				cursor++
				state = a
			}
		case f:
			if !tape[cursor] {
				tape[cursor] = true
				cursor++
				state = a
			} else {
				tape[cursor] = true
				cursor++
				state = e
			}
		default:
			log.Fatalln("Invalid state")
		}
	}

	return len(tape)
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
