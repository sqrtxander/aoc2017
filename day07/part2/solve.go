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

type tree struct {
	weight      int
	totalWeight int
	children    []*tree
}

func (t *tree) getWeights() int {
	result := t.weight
	if t.children != nil {
		for _, child := range t.children {
			if child != nil {
				result += child.getWeights()
			}
		}
	}
	t.totalWeight = result
	return result
}

func (t *tree) hasBalancedChildren() bool {
	if t == nil || t.children == nil || len(t.children) == 0 {
		return true
	}
	toMatch := t.children[0].totalWeight
	return utils.All(t.children, func(c *tree) bool {
		return c.totalWeight == toMatch
	})
}

func (t *tree) findActual() int {
	o := getOddTree(t.children)
	normalWeight := getNormalWeight(t.children)
	if !t.hasBalancedChildren() && o.hasBalancedChildren() {
		return o.weight + normalWeight - o.totalWeight
	}
	return o.findActual()
}

func getNormalWeight(trees []*tree) int {
	if trees == nil {
		return -1
	}
	normal := map[int]bool{}
	for _, t := range trees {
		if t == nil {
			continue
		}
		if _, ok := normal[t.totalWeight]; !ok {
			normal[t.totalWeight] = false
		} else {
			return t.totalWeight
		}
	}
	return -1
}

func getOddTree(trees []*tree) *tree {
	if trees == nil {
		return nil
	}
	unique := map[int]bool{}
	for _, t := range trees {
		if t == nil {
			continue
		}
		if _, ok := unique[t.totalWeight]; !ok {
			unique[t.totalWeight] = true
		} else {
			unique[t.totalWeight] = false
		}
	}
	for _, t := range trees {
		if t == nil {
			continue
		}
		if unique[t.totalWeight] {
			return t
		}
	}
	return nil
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	lines := strings.Split(s, "\n")

	treesAdj := map[utils.Pair[string, int]][]string{}
	for _, line := range lines {
		words := strings.Fields(line)
		name := words[0]
		children := words[min(3, len(words)-1):]
		val := utils.HandledAtoi(words[1])
		treesAdj[utils.Pair[string, int]{K: name, V: val}] = children
	}
	trees := map[string]*tree{}
	for t := range treesAdj {
		trees[t.K] = &tree{weight: t.V}
	}

	for t, tns := range treesAdj {
		for _, tn := range tns {
			trees[t.K].children = append(trees[t.K].children, trees[tn])
		}
	}
	rootFinder := map[string]bool{}
	for t := range treesAdj {
		rootFinder[t.K] = true
	}
	for _, tns := range treesAdj {
		for _, tn := range tns {
			delete(rootFinder, tn)
		}
	}
	var root tree
	for t := range rootFinder {
		root = *trees[t]
	}
	root.getWeights()
	return root.findActual()
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
