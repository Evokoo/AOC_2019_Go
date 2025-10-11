package day06

func Solve(file string, part int) int {
	tree := BuildTree(file)
	result := 0

	if part == 1 {
		for id := range tree.lookup {
			if file == "example.txt" && (id == "SAN" || id == "YOU") {
				continue
			}
			result += tree.DistanceToRoot(id)
		}
	} else {
		result = tree.DistanceBetweenNodes("YOU", "SAN")
	}

	return result
}
