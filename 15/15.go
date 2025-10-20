package day15

func Solve(file string, part int) int {
	program := ParseInput(file)
	zone := NewZone()

	distance := zone.Scan(program)
	time := zone.Flood()

	switch part {
	case 1:
		return distance
	case 2:
		return time
	default:
		panic("Invalid part")
	}
}
