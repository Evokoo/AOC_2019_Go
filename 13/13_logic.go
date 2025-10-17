package day13

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
// GAME
// ========================
type Game struct {
	walls  map[Point]struct{}
	blocks map[Point]struct{}
	paddle Point
	ball   Point
	score  int
}

func NewGame() *Game {
	return &Game{
		walls:  make(map[Point]struct{}),
		blocks: make(map[Point]struct{}),
		paddle: Point{},
		ball:   Point{},
	}
}

func (g *Game) Build(cpu *intcode.CPU) {
	cpu.Run()

	for cpu.HasOutput() {
		id := cpu.ReadOutput()
		y := cpu.ReadOutput()
		x := cpu.ReadOutput()

		if x == -1 && y == 0 {
			g.score = id
		}

		switch id {
		case 0:
			delete(g.blocks, Point{x, y})
		case 1:
			g.AddWall(x, y)
		case 2:
			g.AddBlock(x, y)
		case 3:
			g.SetPaddle(x, y)
		case 4:
			g.SetBall(x, y)
		}
	}
}

func (g *Game) Play(c *intcode.CPU) int {
	c.WriteToMemory(2, 0)

	score := 0
	for c.IsActive() {
		g.Build(c)

		if g.paddle.x < g.ball.x {
			c.ReadInput(1)
		} else if g.paddle.x > g.ball.x {
			c.ReadInput(-1)
		} else {
			c.ReadInput(0)
		}

		if g.score > score {
			score = g.score
		}
	}
	return score
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
func (g *Game) GetBlockCount() int {
	return len(g.blocks)
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
