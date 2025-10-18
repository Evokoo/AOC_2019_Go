package day14

func Solve(file string) int {
	reactions := ParseInput(file)
	ore := 0

	reactions.RequiredOre("FUEL", 1, make(Stock), &ore)

	return ore
}
