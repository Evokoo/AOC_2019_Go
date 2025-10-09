package day01

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// FUEL SUM
// ========================

func CalculateFuelRequirements(file string, part int) int {
	data := utils.ReadFile(file)
	sum := 0

	for amount := range strings.SplitSeq(data, "\n") {
		n, _ := strconv.Atoi(amount)

		if part == 1 {
			sum += n/3 - 2
		} else {
			sum += CalculateAddedFuel(n)
		}
	}
	return sum
}

func CalculateAddedFuel(n int) (sum int) {
	for {
		n = n/3 - 2
		if n < 0 {
			break
		}
		sum += n
	}
	return
}
