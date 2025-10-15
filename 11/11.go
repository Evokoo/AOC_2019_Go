package day11

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string, part int) int {
	cpu := intcode.NewCPU(ParseInput(file))
	bot := NewRobot()

	if part == 2 {
		bot.Paint(1)
	}

	for cpu.IsActive() {
		cpu.ReadInput(bot.Camera())
		cpu.Run()

		bot.Turn(cpu.ReadOutput())
		bot.Paint(cpu.ReadOutput())
		bot.Step()
	}

	if part == 2 {
		//See Console for result
		bot.PrintSquares()
		return -1
	}

	return bot.CountSquares()
}
