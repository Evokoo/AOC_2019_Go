package day13

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string, part int) int {
	program := ParseInput(file)

	cpu := intcode.NewCPU(program)
	game := NewGame()

	if part == 1 {
		game.Build(cpu)
		return game.GetBlockCount()
	} else {
		cpu.Reset(program)
	}

	return game.Play(cpu)
}
