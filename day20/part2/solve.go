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

type point struct {
	p *utils.Point3D
	v *utils.Point3D
	a *utils.Point3D
}

func parsePoint(line string) point {
	parts := strings.Split(line, ", ")
	parts = utils.Map(parts, func(p string) string {
		return p[3 : len(p)-1]
	})
	pp := utils.Map(strings.Split(parts[0], ","), utils.HandledAtoi)
	vv := utils.Map(strings.Split(parts[1], ","), utils.HandledAtoi)
	aa := utils.Map(strings.Split(parts[2], ","), utils.HandledAtoi)
	return point{
		p: &utils.Point3D{
			X: pp[0],
			Y: pp[1],
			Z: pp[2],
		},
		v: &utils.Point3D{
			X: vv[0],
			Y: vv[1],
			Z: vv[2],
		},
		a: &utils.Point3D{
			X: aa[0],
			Y: aa[1],
			Z: aa[2],
		},
	}
}

func (p *point) move() {
	p.v.Add(*p.a)
	p.p.Add(*p.v)
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	points := utils.Map(lines, parsePoint)

	for range 1000 {
		seen := map[utils.Point3D]bool{}
		for _, p := range points {
			p.move()
			if _, ok := seen[*p.p]; ok {
				seen[*p.p] = true
			} else {
				seen[*p.p] = false
			}
		}
		points = utils.Filter(points, func(p point) bool {
			return !seen[*p.p]
		})
	}

	return len(points)
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
