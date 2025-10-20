package day15

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// DROID
// ========================
type Droid struct {
	position Point
	steps    int
	cpu      *intcode.CPU
}

func NewDroid(position Point, steps int, cpu *intcode.CPU) *Droid {
	return &Droid{
		position: position,
		steps:    steps,
		cpu:      cpu,
	}
}

// ========================
// SCAN AREA
// ========================
type Zone struct {
	tiles    Set
	walls    Set
	o2       Point
	distance int
}

func NewZone() Zone {
	return Zone{
		tiles:    make(Set),
		walls:    make(Set),
		o2:       Point{0, 0},
		distance: 0,
	}
}

var DIRECTIONS = map[int]Point{
	1: {0, -1},
	2: {0, 1},
	3: {-1, 0},
	4: {1, 0},
}

func ScanArea(program []int, zone *Zone) {
	droid := NewDroid(Point{0, 0}, 0, intcode.NewCPU(program))
	queue := NewQueue(droid)

	for !queue.IsEmpty() {
		current := queue.Pop()

		for input, direction := range DIRECTIONS {
			position := Point{current.position.x + direction.x, current.position.y + direction.y}

			if zone.tiles.Has(position) || zone.walls.Has(position) {
				continue
			}

			clone := current.cpu.Clone()
			clone.ReadInput(input)
			clone.Run()

			output := clone.ReadOutput()

			switch output {
			case 0:
				zone.walls.Add(position)
			case 1, 2:
				zone.tiles.Add(position)
				queue.Push(NewDroid(position, current.steps+1, clone))

				if output == 2 {
					zone.o2 = position
					zone.distance = current.steps + 1
				}
			}
		}
	}
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

// ========================
// QUEUE
// ========================
type Queue []*Droid

func NewQueue(start *Droid) Queue {
	return []*Droid{start}
}
func (q *Queue) Push(value *Droid) {
	*q = append(*q, value)
}
func (q *Queue) Pop() *Droid {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// ========================
// SET
// ========================
type Set map[Point]struct{}

func (s *Set) Add(tile Point) {
	(*s)[tile] = struct{}{}
}
func (s *Set) Has(tile Point) bool {
	_, found := (*s)[tile]
	return found
}
func (s *Set) Delete(tile Point) {
	delete(*s, tile)
}
