package day17

func Solve(file string) int {
	scaffold := make(Scaffold)
	scaffold.Build(ParseInput(file))

	return scaffold.AlignmentScore()
}
