package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ",", "", -1)
	lines := strings.Split(s, "\n")

	treesAdj := map[string][]string{}
	for _, line := range lines {
		if !strings.Contains(line, "->") {
			continue
		}
		words := strings.Fields(line)
		data := words[0]
		children := words[3:]
		treesAdj[data] = children
	}
	trees := map[string]bool{}
	for t := range treesAdj {
		trees[t] = true
	}
	for _, tns := range treesAdj {
		for _, tn := range tns {
			delete(trees, tn)
		}
	}
	for t := range trees {
		return t
	}
	return ""
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
