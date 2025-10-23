package day16

func Solve(file string) int {

	sequence := ParseInput(file)
	sequence.Process(100)

	return sequence.Result()
}
