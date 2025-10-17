package day13

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string, part int) int {
	cpu := intcode.NewCPU(ParseInput(file))
	cpu.Run()
	game := BuildGame(cpu.DumpOutput())
	return len(game.blocks)
}
