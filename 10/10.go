package day10

func Solve(file string, part int) int {
	grid := ParseInput(file)
	grid.ScanAstroids()

	if part == 2 {
		return grid.DestoryAstroids(200)
	}

	return grid.visibility
}
