package day25

func Solve(file string) int {
	program := ParseInput(file)
	return FindPassword(program)
}
