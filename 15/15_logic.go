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
	tiles Set
	walls Set
	o2    Point
}

func NewZone() Zone {
	return Zone{
		tiles: make(Set),
		walls: make(Set),
		o2:    Point{0, 0},
	}
}

var DIRECTIONS = map[int]Point{
	1: {0, -1},
	2: {0, 1},
	3: {-1, 0},
	4: {1, 0},
}

func (z *Zone) Scan(program []int) int {
	droid := NewDroid(Point{0, 0}, 0, intcode.NewCPU(program))
	queue := NewQueue(droid)
	distance := 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		for input, direction := range DIRECTIONS {
			position := Point{current.position.x + direction.x, current.position.y + direction.y}

			if z.tiles.Has(position) || z.walls.Has(position) {
				continue
			}

			clone := current.cpu.Clone()
			clone.ReadInput(input)
			clone.Run()

			switch output := clone.ReadOutput(); output {
			case 0:
				z.walls.Add(position)
			case 1, 2:
				z.tiles.Add(position)
				queue.Push(NewDroid(position, current.steps+1, clone))

				if output == 2 {
					z.o2 = position
					distance = current.steps + 1
				}
			}
		}
	}

	return distance
}

func (z *Zone) Flood() int {
	queue := NewQueue(z.o2)
	time := 0

	for !queue.IsEmpty() {
		points := queue.Length()

		for range points {
			current := queue.Pop()

			for _, dir := range DIRECTIONS {
				next := Point{current.x + dir.x, current.y + dir.y}

				if z.tiles.Has(next) {
					z.tiles.Delete(next)
					queue.Push(next)
				}
			}
		}

		if !queue.IsEmpty() {
			time++
		}
	}

	return time
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
type Queue[T any] []T

func NewQueue[T any](start T) Queue[T] {
	return []T{start}
}

func (q *Queue[T]) Push(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Pop() T {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}

func (q *Queue[T]) IsEmpty() bool {
	return (*q).Length() == 0
}
func (q *Queue[T]) Length() int {
	return len(*q)
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
func (s *Set) IsEmpty() bool {
	return len(*s) == 0
}
