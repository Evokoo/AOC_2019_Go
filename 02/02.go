package day02

import "github.com/Evokoo/AOC_2019_Go/intcode"

func Solve(file string, part int) int {
	program := ParseInput(file)
	cpu := intcode.NewCPU(program)

	if part == 2 {
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				cpu.WriteMemory(a, 1) // noun
				cpu.WriteMemory(b, 2) // verb
				cpu.Run()

				if cpu.ReadMemory(0) == 19690720 {
					return 100*a + b
				}

				cpu.Reset(program)
			}
		}
	}

	cpu.WriteMemory(12, 1)
	cpu.WriteMemory(2, 2)
	cpu.Run()

	return cpu.ReadMemory(0)
}
