package day05

import (
	"github.com/Evokoo/AOC_2019_Go/intcode"
)

func Solve(file string, part int) int {
	cpu := intcode.NewCPU(ParseInput(file))
	id := 1

	if part == 2 {
		id = 5
	}

	cpu.PushToInput(id)
	cpu.Run()
	output := cpu.GetOutput()

	return output[len(output)-1]
}
