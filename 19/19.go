package day19

func Solve(file string, part int) int {
	program := ParseInput(file)
	scanner := NewScanner(program)

	switch part {
	case 1:
		return scanner.Run(50)
	case 2:
		return scanner.LocateObject(100)
	}

	return -1
}
