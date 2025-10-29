package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// DROID
// ========================
type Droid struct {
	cpu *intcode.CPU
}

func NewDroid(program []int) Droid {
	return Droid{cpu: intcode.NewCPU(program)}
}

func (d *Droid) RunScript(script []string) {
	d.cpu.Run()

	for _, r := range strings.Join(script, "\n") {
		d.cpu.ReadInput(int(r))
	}

	d.cpu.Run()
}

func (d *Droid) Status(debug bool) (bool, int) {
	output := d.cpu.DumpOutput()

	for _, n := range output {
		if n > 255 {
			return true, n
		}

		if debug {
			char := fmt.Sprintf("%c", n)
			fmt.Print(char)
		}
	}

	return false, 0
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)

	var program []int
	for value := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(value)
		program = append(program, n)
	}

	return program
}
