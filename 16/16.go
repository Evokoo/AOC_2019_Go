package day16

func Solve(file string, part int) int {
	sequence := ParseInput(file)

	switch part {
	case 1:
		sequence.Process(100)
	case 2:
		sequence.ProcessReal(sequence.ExtractSlice(0, 7), 100)
	}

	return sequence.ExtractSlice(0, 8)
}
