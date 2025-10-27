package day19

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// SCANNER
// ========================

type Scanner struct {
	cpu *intcode.CPU
}

func NewScanner(program []int) Scanner {
	return Scanner{intcode.NewCPU(program)}
}
func (s *Scanner) CheckPoint(x, y int) int {
	clone := s.cpu.Clone()
	clone.ReadInput(x)
	clone.ReadInput(y)
	clone.Run()
	return clone.ReadOutput()
}

func (s *Scanner) Run(size int) (count int) {
	for x := range size {
		for y := range size {
			if result := s.CheckPoint(x, y); result == 1 {
				count++
			}
		}
	}
	return
}

func (s *Scanner) LocateObject(size int) int {
	vector := PredictBeamVector(*s)
	slope := vector.x / vector.y
	y := size - 1

	for {
		xPred := int(slope * float64(y))

		for x := xPred - 5; x <= xPred+5; x++ {
			topRightX := x + size - 1
			topRightY := y - size + 1

			if (*s).CheckPoint(x, y) == 1 && (*s).CheckPoint(topRightX, topRightY) == 1 {
				return x*10_000 + topRightY
			}
		}
		y++
	}
}

// ========================
// POINT & VECTOR
// ========================
type Point struct{ x, y int }
type Vector struct{ x, y float64 }

// ========================
// SLOPE
// ========================
func PredictBeamVector(scanner Scanner) Vector {
	points := []Point{{0, 0}}

	for _, y := range []int{100, 300, 500, 700} {
		x := 0
		for scanner.CheckPoint(x, y) != 1 {
			x++
		}
		points = append(points, Point{x, y})
	}

	a := points[0]
	b := points[len(points)-1]

	return Vector{float64(b.x - a.x), float64(b.y - a.y)}
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
