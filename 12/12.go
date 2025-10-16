package day12

func Solve(file string) int {
	moons := ParseInput(file)
	moons.Simulate(1000)

	return moons.CalcuateTotalEnergy()
}
