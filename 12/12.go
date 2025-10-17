package day12

func Solve(file string, part int) int {
	moons := ParseInput(file)

	if part == 2 {
		return moons.PredictLoop()
	}

	moons.Simulate(1000)
	return moons.CalcuateTotalEnergy()
}
