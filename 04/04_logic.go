package day04

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// RANGE
// ========================
type Range [2]int

// ========================
// VALIDATOR
// ========================

func CountValidNumbers(valueRange Range, part int) (count int) {
	for i := valueRange[0]; i <= valueRange[1]; i++ {
		if IsValid(i, part) {
			count++
		}
	}
	return
}
func IsValid(n int, part int) bool {
	double := false
	prev := n % 10
	consecutive := 1
	n /= 10

	for n > 0 {
		next := n % 10

		if next > prev {
			return false
		}

		if next == prev {
			consecutive++
		} else {
			if part == 1 && consecutive >= 2 {
				double = true
			}
			if part == 2 && consecutive == 2 {
				double = true
			}
			consecutive = 1
		}

		prev = next
		n /= 10
	}

	if part == 1 && consecutive >= 2 {
		double = true
	}
	if part == 2 && consecutive == 2 {
		double = true
	}

	return double
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Range {
	data := utils.ReadFile(file)
	var output Range
	for i, value := range strings.Split(data, "-") {
		n, _ := strconv.Atoi(value)
		output[i] = n
	}
	return output
}
