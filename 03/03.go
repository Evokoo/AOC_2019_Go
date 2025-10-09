package day03

func Solve(file string, part int) int {
	paths := ParseInput(file)
	wires := []*Wire{NewWire(), NewWire()}

	for i, path := range paths {
		for _, move := range path {
			wires[i].UpdateWire(move)
		}
	}

	return FindClosestIntersection(wires, part)
}
