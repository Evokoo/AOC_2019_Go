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
// AMPLIFERS
// ========================
type Amplifers []*intcode.CPU

func SetupAmplifiers(program, sequence []int) Amplifers {
	var amplifiers Amplifers

	for i := range 5 {
		amp := intcode.NewCPU(program)
		amp.PushToInput(sequence[i])
		amplifiers = append(amplifiers, amp)
	}

	return amplifiers
}

func RunSequence(amplifiers Amplifers, part int) int {
	signal := 0

	if part == 1 {
		for _, amp := range amplifiers {
			amp.PushToInput(signal)
			amp.Run()
			signal = amp.PopFromOutput()
		}
		return signal
	}

	for amplifiers[4].IsActive() {
		for _, amp := range amplifiers {
			amp.PushToInput(signal)
			amp.Run()
			signal = amp.PopFromOutput()
		}
	}

	return signal
}

func FindOptimalSignal(file string, part int) int {
	program := ParseInput(file)
	maxSignal := 0
	for _, sequence := range GenerateSequences(part) {
		signal := RunSequence(SetupAmplifiers(program, sequence), part)
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	return maxSignal
}

func GenerateSequences(part int) [][]int {
	phaseSet := []int{0, 1, 2, 3, 4}

	if part == 2 {
		phaseSet = []int{5, 6, 7, 8, 9}
	}

	var sequences [][]int
	permutate(phaseSet, 0, &sequences)

	return sequences
}

func permutate(a []int, l int, results *[][]int) {
	if l == len(a)-1 {
		tmp := make([]int, len(a))
		copy(tmp, a)
		*results = append(*results, tmp)
		return
	}

	for i := l; i < len(a); i++ {
		a[l], a[i] = a[i], a[l]
		permutate(a, l+1, results)
		a[l], a[i] = a[i], a[l]
	}
}
