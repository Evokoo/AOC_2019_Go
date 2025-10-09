package day02

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

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
