package day17

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// Point
// ========================
type Point struct{ x, y int }

var DIRECTIONS = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// ========================
// Scaffold
// ========================
type Scaffold map[Point]struct{}

func (s *Scaffold) Build(program []int) {
	cpu := intcode.NewCPU(program)
	cpu.Run()

	x, y := 0, 0
	for _, value := range cpu.DumpOutput() {
		switch value {
		case 35:
			(*s).Add(x, y)
			x++
		case 46:
			x++
		case 10:
			y++
			x = 0
		}
	}
}

func (s *Scaffold) Add(x, y int) {
	(*s)[Point{x, y}] = struct{}{}
}

func (s *Scaffold) Has(p Point) bool {
	_, found := (*s)[p]
	return found
}

func (s *Scaffold) AlignmentScore() int {
	score := 0

points:
	for point := range *s {
		for _, d := range DIRECTIONS {
			next := Point{point.x + d.x, point.y + d.y}
			if !(*s).Has(next) {
				continue points
			}
		}

		score += point.x * point.y
	}

	return score
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)

	var program []int
	for value := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(value)
		program = append(program, n)
	}

	return program
}
