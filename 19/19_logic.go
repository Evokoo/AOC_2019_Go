package day19

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// SCANNER
// ========================

func RunScan(size int, program []int) (count int) {
	cpu := intcode.NewCPU(program)

	for x := range size {
		for y := range size {
			fmt.Println(x, y)
			cpu.ReadInput(x)
			cpu.ReadInput(y)
			cpu.Run()

			switch cpu.ReadOutput() {
			case 1:
				count++
			}

			cpu.Reset(program)
		}
	}
	return
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)

	var program []int
	for value := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(value)
		program = append(program, n)
	}

	return program
}
