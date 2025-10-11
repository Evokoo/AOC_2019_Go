package day06

import "fmt"

func Solve(file string) int {
	tree := BuildTree(file)
	orbits := 0

	for key, node := range tree.lookup {
		if key == "COM" {
			continue
		}
		steps := 0

		for node.id != tree.root.id {
			node = node.parent
			steps++
		}

		orbits += steps
	}

	fmt.Println(orbits)
	return orbits
}
