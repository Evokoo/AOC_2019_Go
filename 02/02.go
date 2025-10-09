package day02

func Solve(file string) int {
	program := ParseInput(file)

	if file == "input.txt" {
		program.SetValue(12, 1)
		program.SetValue(2, 2)
	}

	program.Run()

	return program.GetValue(0)
}
