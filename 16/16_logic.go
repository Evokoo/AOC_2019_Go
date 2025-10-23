package day16

import (
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// SEQUENCE
// ========================
type Sequence []int

func (s *Sequence) Process(rounds int) {
	length := len(*s)
	next := make(Sequence, length)

	for range rounds {

		for i := range length {
			value := 0
			for j := range length {
				value += (*s)[j] * GetPatternValue(i, j)
			}
			next[i] = ExtactLastDigit(value)
		}
		*s = next
	}
}

func (s *Sequence) ProcessReal(offset, rounds int) {
	length := len(*s)
	totalLen := length * 10_000

	keepLen := totalLen - offset
	signal := make([]int, keepLen)

	for i := 0; i < keepLen; i++ {
		orig := (offset + i) % length
		signal[i] = (*s)[orig]
	}

	for p := 0; p < rounds; p++ {
		sum := 0
		for i := keepLen - 1; i >= 0; i-- {
			sum = (sum + signal[i]) % 10
			signal[i] = sum
		}
	}

	*s = signal
}

func (s *Sequence) ExtractSlice(start, size int) int {
	num := 0
	for _, d := range (*s)[start : start+size] {
		num = num*10 + d
	}
	return num
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Sequence {
	data := utils.ReadFile(file)
	output := make(Sequence, 0, len(data))

	for _, r := range data {
		output = append(output, int(r-'0'))
	}
	return output
}

func ExtactLastDigit(n int) int {
	return utils.Abs(n) % 10
}

var Pattern = []int{0, 1, 0, -1}

func GetPatternValue(i, j int) int {
	return Pattern[((j+1)/(i+1))%4]
}
