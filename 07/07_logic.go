package day07

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)

	var program []int
	for code := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(code)
		program = append(program, n)
	}

	return program
}

// ========================
// THRUSTERS
// ========================

func RunThrusterSequence(program, sequence []int) int {
	cpu := intcode.NewCPU(program)
	signal := 0

	for i := 0; i < 5; i++ {
		cpu.PushToInput(sequence[i])
		cpu.PushToInput(signal)
		cpu.Run()
		signal = cpu.GetOutput()[0]
		cpu.Reset(program)
	}

	return signal
}

func CalibrateThrusters(file string) int {
	program := ParseInput(file)

	var sequences [][]int
	GenerateSequences([]int{0, 1, 2, 3, 4}, 0, &sequences)

	var max int
	for _, sequence := range sequences {
		signal := RunThrusterSequence(program, sequence)
		if signal > max {
			max = signal
		}
	}
	return max
}

func GenerateSequences(a []int, l int, results *[][]int) {
	if l == len(a)-1 {
		tmp := make([]int, len(a))
		copy(tmp, a)
		*results = append(*results, tmp)
		return
	}

	for i := l; i < len(a); i++ {
		a[l], a[i] = a[i], a[l]
		GenerateSequences(a, l+1, results)
		a[l], a[i] = a[i], a[l]
	}
}
