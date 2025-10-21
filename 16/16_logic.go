package day16

import (
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// SEQUENCE
// ========================
type Sequence []int

var Pattern = []int{0, 1, 0, -1}

func (s *Sequence) Process(rounds int){
	next := make(Sequence, 0, len(*s))

	for
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
