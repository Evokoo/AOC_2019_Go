package day03

import (
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// WIRE
// ========================
type Wire struct {
	position Point
	path     map[Point]int
	steps    int
}

func NewWire() *Wire {
	return &Wire{
		position: Point{0, 0},
		path:     make(map[Point]int),
		steps:    0,
	}
}

func (w *Wire) UpdateWire(move Move) {
	for i := 0; i < move.steps; i++ {
		w.position.x += move.direction.x
		w.position.y += move.direction.y
		w.steps++
		w.path[w.position] = w.steps
	}
}

// ========================
// MOVE
// ========================
type Move struct {
	direction Point
	steps     int
}

// ========================
// PARSER
// ========================

func ParseInput(file string) [][]Move {
	data := utils.ReadFile(file)

	var paths [][]Move

	for path := range strings.SplitSeq(data, "\n") {
		var sequence []Move

		for input := range strings.SplitSeq(path, ",") {
			var move Move

			switch input[0] {
			case 'R':
				move.direction = Point{1, 0}
			case 'L':
				move.direction = Point{-1, 0}
			case 'U':
				move.direction = Point{0, 1}
			case 'D':
				move.direction = Point{0, -1}
			}

			steps, _ := strconv.Atoi(input[1:])
			move.steps = steps
			sequence = append(sequence, move)
		}
		paths = append(paths, sequence)
	}

	return paths
}

// ========================
// DISTANCE
// ========================
func ManhattanDistance(a Point, b Point) int {
	return utils.Abs(a.x-b.x) + utils.Abs(a.y-b.y)
}

func FindClosestIntersection(wires []*Wire, part int) int {
	intersections := make(map[Point]int)

	for point := range wires[0].path {
		if _, has := wires[1].path[point]; has {
			intersections[point] = wires[0].path[point] + wires[1].path[point]
		}
	}

	minDistance := math.MaxInt
	minSteps := math.MaxInt
	origin := Point{0, 0}

	for point, steps := range intersections {
		distance := ManhattanDistance(origin, point)
		if distance < minDistance {
			minDistance = distance
		}

		if steps < minSteps {
			minSteps = steps
		}
	}

	if part == 2 {
		return minSteps
	}

	return minDistance
}
