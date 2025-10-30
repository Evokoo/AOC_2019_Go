package day25

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

func (d *Droid) Clone() Droid {
	return Droid{cpu: d.cpu.Clone()}
}

func (d *Droid) ReadOutput() string {
	var output strings.Builder

	for _, n := range d.cpu.DumpOutput() {
		output.WriteString(fmt.Sprintf("%c", n))
	}

	return output.String()
}

func (d *Droid) RunCommand(commands []string) {
	for _, command := range commands {
		var s string
		switch command {
		case "E":
			s = "east\n"
		case "W":
			s = "west\n"
		case "S":
			s = "south\n"
		case "N":
			s = "north\n"
		case "T":
			s = "take "
		case "D":
			s = "drop "
		case "I":
			s = "inv\n"
		default:
			s = command + "\n"
		}

		for _, r := range s {
			d.cpu.ReadInput(int(r))
		}
	}

	d.cpu.Run()
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

// ========================
// FIND WEIGHT
// ========================

var collectItems = []string{
	"N", "N", "N", "T", "mutex", //mutex
	"S", "S", "E", "N", "T", "loom", //loom
	"S", "W", "S", "E", "T", "semiconductor", //semiconductor
	"E", "T", "ornament", //ornament
	"W", "W", "W", "W", "T", "sand", //sand
	"S", "E", "T", "asterisk", //asterisk
	"N", "T", "wreath", //wreath
	"S", "W", "N", "N", "T", "dark matter", //dark matter
	"E", //
	"D", "dark matter", "D", "wreath", "D", "asterisk", "D", "sand", "D", "ornament",
	"D", "semiconductor", "D", "loom", "D", "mutex", // Drop all items
}

var items = []string{"mutex", "loom", "semiconductor", "ornament", "sand", "asterisk", "wreath", "dark matter"}

func FindPassword(program []int) int {
	original := NewDroid(program)
	original.RunCommand(collectItems)
	original.cpu.ClearOutput()

	for mask := 0; mask < (1 << len(items)); mask++ {
		clone := original.Clone()
		var command []string

		for i := 0; i < len(items); i++ {
			if (mask>>i)&1 == 1 {
				command = append(command, "T")
				command = append(command, items[i])
			}
		}

		//Final Step
		command = append(command, "E")
		clone.RunCommand(command)

		response := clone.ReadOutput()

		if !strings.Contains(response, "ejected") {
			match := utils.QuickMatch(response, `\d+`)[0]
			password, _ := strconv.Atoi(match)
			return password
		}
	}

	panic("Password Not Found")
}
