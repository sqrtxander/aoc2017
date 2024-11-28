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

func solve(s string, l int) int {
	s = strings.TrimSpace(s)
	lengths := utils.Map(strings.Split(s, ","), utils.HandledAtoi)
	skip := byte(0)
	pos := byte(0)
	nums := make([]byte, l)
	for i := range l {
		nums[i] = byte(i)
	}

	for _, length := range lengths {
		// reversing
		for i, j := int(pos), int(pos)+int(length)-1; i < j; i, j = i+1, j-1 {
			nums[i%l], nums[j%l] = nums[j%l], nums[i%l]
		}
		pos = pos + byte(length) + skip
		skip++
	}

	return int(nums[0]) * int(nums[1])
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
	fmt.Println(solve(string(contents), 256))
}
