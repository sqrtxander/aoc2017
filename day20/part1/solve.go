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

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	points := utils.Map(lines, parsePoint)
	largeT := 1000000
	largeTDists := utils.Map(points, func(p point) int {
		p.a.X *= largeT * largeT / 2
		p.a.Y *= largeT * largeT / 2
		p.a.Z *= largeT * largeT / 2
		p.v.X *= largeT
		p.v.Y *= largeT
		p.v.Z *= largeT
		p.p.Add(*p.v)
		p.p.Add(*p.a)
		return p.p.Manhattan3D()
	})
	minDist := slices.Min(largeTDists)
	return slices.Index(largeTDists, minDist)
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
