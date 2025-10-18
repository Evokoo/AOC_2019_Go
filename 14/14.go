package day14

func Solve(file string, part int) int {
	reactions := ParseInput(file)
	ore := 0

	if part == 1 {
		reactions.RequiredOre("FUEL", 1, make(Stock), &ore)
		return ore
	}

	limit := 1000000000000
	low, high := 1, limit
	fuel := 0

	for low <= high {
		mid := (low + high) / 2
		ore := 0
		reactions.RequiredOre("FUEL", mid, make(Stock), &ore)

		if ore <= limit {
			fuel = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return fuel
}
