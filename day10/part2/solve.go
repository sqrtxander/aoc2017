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
	l := 256
	lengths := []byte(s)
	lengths = append(lengths,
		17,
		31,
		73,
		47,
		23,
	)
	skip := byte(0)
	pos := byte(0)
	nums := make([]byte, l)
	for i := range l {
		nums[i] = byte(i)
	}

	for range 64 {
		for _, length := range lengths {
			// reversing
			for i, j := int(pos), int(pos)+int(length)-1; i < j; i, j = i+1, j-1 {
				nums[i%l], nums[j%l] = nums[j%l], nums[i%l]
			}
			pos = pos + byte(length) + skip
			skip++
		}
	}
	result := ""
	contribution := byte(0)
	for i, num := range nums {
		contribution ^= num
		if i%16 == 15 {
			result += fmt.Sprintf("%02x", contribution)
			contribution = 0
		}
	}
	return result
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
