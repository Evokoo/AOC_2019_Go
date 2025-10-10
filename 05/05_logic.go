package day05

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// PARSER
// ========================
type Program []int

func ParseInput(file string) Program {
	data := utils.ReadFile(file)
	var program Program
	for value := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(value)
		program = append(program, n)
	}
	return program
}
