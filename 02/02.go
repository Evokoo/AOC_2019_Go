package day02

import "github.com/Evokoo/AOC_2019_Go/intcode"

func Solve(file string, part int) int {
	program := ParseInput(file)
	cpu := intcode.NewCPU(program)

	if part == 2 {
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				cpu.WriteToMemory(a, 1) // noun
				cpu.WriteToMemory(b, 2) // verb
				cpu.Run()

				if cpu.ReadFromMemory(0) == 19690720 {
					return 100*a + b
				}

				cpu.Reset(program)
			}
		}
	}

	cpu.WriteToMemory(12, 1)
	cpu.WriteToMemory(2, 2)
	cpu.Run()

	return cpu.ReadFromMemory(0)
}
