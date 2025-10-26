package day17

func Solve(file string, part int) int {
	scaffold := make(Scaffold)

	switch part {
	case 1:
		scaffold.Build(ParseInput(file))
		return scaffold.AlignmentScore()
	case 2:
		return 0
	default:
		panic("Invalid Part")
	}

}
