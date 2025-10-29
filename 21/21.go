package day21

func Solve(file string, part int) int {
	program := ParseInput(file)
	droid := NewDroid(program)

	switch part {
	case 1:
		script := []string{
			"NOT C J",
			"AND D J",
			"NOT A T",
			"OR T J",
			"WALK\n",
		}
		droid.RunScript(script)
	case 2:
		script := []string{
			"NOT C J",
			"AND D J",
			"AND H J",
			"NOT B T",
			"AND D T",
			"OR T J",
			"NOT A T",
			"OR T J",
			"RUN\n",
		}
		droid.RunScript(script)
	}

	if success, result := droid.Status(false); success {
		return result
	}

	return 0
}
