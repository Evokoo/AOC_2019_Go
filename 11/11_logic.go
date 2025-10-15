package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// ROBOT
// ========================
type Robot struct {
	position   Point
	bearing    int
	squares    map[Point]int
	dimensions [4]int
}

func NewRobot() *Robot {
	return &Robot{
		position: Point{0, 0},
		bearing:  0,
		squares:  make(map[Point]int),
	}
}
func (r *Robot) Paint(colour int) {
	r.squares[r.position] = colour
}
func (r *Robot) Turn(direction int) {
	degrees := 90
	if direction == 0 {
		degrees = -90
	}
	r.bearing = (r.bearing + degrees + 360) % 360
}
func (r *Robot) Step() {
	switch r.bearing {
	case 0:
		r.position.y--
	case 90:
		r.position.x++
	case 180:
		r.position.y++
	case 270:
		r.position.x--
	}

	if r.position.x < r.dimensions[0] {
		r.dimensions[0] = r.position.x
	}
	if r.position.y < r.dimensions[1] {
		r.dimensions[1] = r.position.y
	}
	if r.position.x > r.dimensions[2] {
		r.dimensions[2] = r.position.x
	}
	if r.position.y > r.dimensions[3] {
		r.dimensions[3] = r.position.y
	}
}
func (r *Robot) Camera() int {
	if colour, found := r.squares[r.position]; found {
		return colour
	}
	return 0
}
func (r *Robot) CountSquares() int {
	return len(r.squares)
}
func (r *Robot) PrintSquares() {
	cols := r.dimensions[2] - r.dimensions[0] + 1
	rows := r.dimensions[3] - r.dimensions[1] + 1

	var row strings.Builder

	for y := range rows {
		for x := range cols {
			switch r.squares[Point{x, y}] {
			case 0:
				row.WriteRune('░')
			case 1:
				row.WriteRune('▓')
			}
		}
		fmt.Println(row.String())
		row.Reset()
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
