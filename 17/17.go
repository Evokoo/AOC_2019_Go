package day17

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string) int {
	program := ParseInput(file)
	cpu := intcode.NewCPU(program)
	cpu.Run()

	scaffold := make(Scaffold)
	scaffold.Build(cpu.DumpOutput())

	return scaffold.AlignmentScore()
}
