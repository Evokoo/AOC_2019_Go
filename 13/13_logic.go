package day13

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// GAME
// ========================
type Game struct {
	walls  map[Point]struct{}
	blocks map[Point]struct{}
	paddle Point
	ball   Point
}

func NewGame() *Game {
	return &Game{
		walls:  make(map[Point]struct{}),
		blocks: make(map[Point]struct{}),
		paddle: Point{},
		ball:   Point{},
	}
}

func BuildGame(data []int) *Game {
	game := NewGame()

	for i := 0; i < len(data); i = i + 3 {
		instruction := data[i : i+3]

		switch instruction[2] {
		case 1:
			game.AddWall(instruction[0], instruction[1])
		case 2:
			game.AddBlock(instruction[0], instruction[1])
		case 3:
			game.SetPaddle(instruction[0], instruction[1])
		case 4:
			game.SetBall(instruction[0], instruction[1])
		}
	}

	return game
}

func (g *Game) AddWall(x, y int) {
	g.walls[Point{x, y}] = struct{}{}
}
func (g *Game) AddBlock(x, y int) {
	g.blocks[Point{x, y}] = struct{}{}
}
func (g *Game) SetBall(x, y int) {
	g.ball = Point{x, y}
}
func (g *Game) SetPaddle(x, y int) {
	g.paddle = Point{x, y}
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
