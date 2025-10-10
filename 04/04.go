package day04

func Solve(file string, part int) int {
	valueRange := ParseInput(file)
	return CountValidNumbers(valueRange, part)
}
