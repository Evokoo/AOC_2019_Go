package day02

import "github.com/Evokoo/AOC_2019_Go/intcode"

func Solve(file string, part int) int {
	program := ParseInput(file)
	cpu := intcode.NewCPU(program)

	if part == 2 {
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				cpu.SetValue(a, 1) // noun
				cpu.SetValue(b, 2) // verb
				cpu.Run()

				if cpu.GetValue(0) == 19690720 {
					return 100*a + b
				}

				cpu.Reset(program)
			}
		}
	}

	cpu.SetValue(12, 1)
	cpu.SetValue(2, 2)
	cpu.Run()

	return cpu.GetValue(0)
}
