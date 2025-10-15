package day08

func Solve(file string, part int) int {
	if part == 2 {
		CombineLayers(file, 25, 6)
		//See console for solution
		return -1
	}

	return CheckLayers(file, 25, 6)
}
