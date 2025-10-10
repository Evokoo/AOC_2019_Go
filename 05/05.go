package day05

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string) int {
	program := ParseInput(file)
	cpu := intcode.NewCPU(program)

	cpu.WriteToInput(1)
	cpu.Run()
	output := cpu.GetOutput()

	return output[len(output)-1]
}
