package day09

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string, part int) int {
	cpu := intcode.NewCPU(ParseInput(file))
	cpu.ReadInput(part)
	cpu.Run()

	return cpu.ReadOutput()
}
